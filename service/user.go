package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
	"myblog/config"
	"myblog/dao"
	"myblog/model"
	"myblog/router/middleware"
	"myblog/util"
	"myblog/util/r"
	"sort"
	"strconv"
	"strings"
	"time"
)

type UserService struct {
}

func (*UserService) Login(c *gin.Context, username string, password string) (model.UserInfoDTO, int) {
	userDetailDTO, code := LoadUserByUsername(c, username)
	if code != r.SUCCESS {
		return model.UserInfoDTO{}, code
	}
	if userDetailDTO.Password != util.Generator.MD5(password) {
		return model.UserInfoDTO{}, r.PasswordWrong
	}
	userInfoDTO := userDetailDTO.UserInfoDTO
	// 这里写的比较简单，每个用户只设定一个角色，所以只需要获取第一个值就行了
	//如果想要设置多个角色，就往Token里存数组就行
	uuid := util.Generator.MD5(
		userDetailDTO.IpAddress + userDetailDTO.Browser + userDetailDTO.Os,
	)
	token, code := middleware.CreateToken(userDetailDTO.UserInfoId, userDetailDTO.RoleList[0], uuid)
	if code != r.SUCCESS {
		return model.UserInfoDTO{}, code
	}
	updateLogin(userInfoDTO)
	userInfoDTO.Token = token

	session := sessions.Default(c)
	sessionInfo := model.SessionInfo{UserDetailDTO: userDetailDTO}
	session.Set(User+uuid, sessionInfo)
	rdb.Set(User+uuid, util.Json.Marshal(sessionInfo), time.Duration(config.SessionConf.MaxAge)*time.Second)

	// 下面这个是为了博客用户上传头像用的，因为前端那块不知道怎么取携带Authorization，其他地方都携带上了
	session.Set("userInfoId", userInfoDTO.UserInfoId)

	session.Save()

	return userInfoDTO, code
}

func (*UserService) GetUserAreas(condition model.Condition) (list []model.UserAreaDTO) {
	switch condition.Type {
	case 1: // 用户
		userAreaJSON := rdb.Get(UserArea).Val()
		if userAreaJSON != "" {
			util.Json.Unmarshal(userAreaJSON, &list)
		}
		return
	case 2: // 游客
		visitorAreaMap := rdb.HGetAll(VisitorArea).Val()
		if visitorAreaMap != nil {
			for k, v := range visitorAreaMap {
				value, _ := strconv.Atoi(v)
				dto := model.UserAreaDTO{
					Name:  k,
					Value: value,
				}
				list = append(list, dto)
			}
		}
		return
	default:
		break
	}
	return
}

// 统计地域分布定时任务
func (*UserService) statisticalUserArea() {
	userAreaList := dao.List([]model.UserAuth{}, "ip_source", "", "")
	userAreaMap := make(map[string]int)
	for _, userArea := range userAreaList {
		idx := strings.Index(userArea.IpSource, "省")
		if idx == -1 {
			userAreaMap[userArea.IpSource]++
		} else {
			address := strings.Split(userArea.IpSource, "省")[0]
			userAreaMap[address]++
		}
	}
	var list []model.UserAreaDTO
	for k, v := range userAreaMap {
		dto := model.UserAreaDTO{
			Name:  k,
			Value: v,
		}
		list = append(list, dto)
	}
	userAreaJSON, _ := json.Marshal(list)
	rdb.Set(UserArea, userAreaJSON, 0)
}

func (*UserService) GetList(condition model.Condition) model.PageResult[[]model.BackendUserDTO] {
	count := userDao.CountUser(condition)
	if count == 0 {
		return model.PageResult[[]model.BackendUserDTO]{count, []model.BackendUserDTO{}}
	}
	userList := userDao.UserList(condition)
	return model.PageResult[[]model.BackendUserDTO]{count, userList}
}

func (*UserService) UpdateRole(data model.UserRoleVO) {
	userInfo := model.UserInfo{
		Universal: model.Universal{ID: data.UserInfoId},
		Nickname:  data.Nickname,
	}
	dao.Updates(&userInfo)
	dao.Delete(model.UserRole{}, "user_id = ?", data.UserInfoId)
	var userRoleList []model.UserRole
	for _, id := range data.RoleIdList {
		userRoleList = append(userRoleList, model.UserRole{
			RoleId: id,
			UserId: data.UserInfoId,
		})
	}
	dao.Create(&userRoleList)
}

func (*UserService) UpdateDisable(data model.UserInfo) {
	dao.Updates(&data, "is_disable")
}

func (*UserService) UserOnlineList() model.PageResult[[]model.UserOnlineDTO] {
	keys := rdb.Keys(User + "*").Val()
	count := len(keys)
	if len(keys) == 0 {
		return model.PageResult[[]model.UserOnlineDTO]{}
	}
	var onlineList []model.UserOnlineDTO
	for _, key := range keys {
		var sessionInfo model.SessionInfo
		util.Json.Unmarshal(rdb.Get(key).Val(), &sessionInfo)
		onlineList = append(onlineList, model.UserOnlineDTO{
			UserIndoId:    sessionInfo.UserInfoId,
			Nickname:      sessionInfo.Nickname,
			Avatar:        sessionInfo.Avatar,
			IpAddress:     sessionInfo.IpAddress,
			IpSource:      sessionInfo.IpSource,
			Browser:       sessionInfo.Browser,
			OS:            sessionInfo.Os,
			LastLoginTime: sessionInfo.LastLoginTime,
		})
	}
	sort.Slice(onlineList, func(i, j int) bool {
		return onlineList[i].LastLoginTime.Unix() > onlineList[j].LastLoginTime.Unix()
	})
	return model.PageResult[[]model.UserOnlineDTO]{count, onlineList}
}

func (*UserService) ForceOffline(userInfo model.UserDetailDTO) {
	key := User + util.Generator.MD5(userInfo.IpAddress+userInfo.Browser+userInfo.Os)
	var sessionInfo model.SessionInfo
	util.Json.Unmarshal(rdb.Get(key).Val(), &sessionInfo)
	sessionInfo.IsOffline = 1
	rdb.Del(key)
	rdb.Set(Delete+key[5:], util.Json.Marshal(sessionInfo), time.Duration(config.GinConf.ExpireTime)*time.Hour)
}

func (*UserService) UpdateInfo(userInfo model.UserInfo) {
	dao.Updates(&userInfo, "nickname", "intro", "web_site")
}

func (*UserService) UpdateAvatar(userInfo model.UserInfo) {
	dao.Updates(&userInfo, "avatar")
}

func (*UserService) UpdateAdminPassword(pwdVO model.PasswordVO, id int) int {
	user := dao.GetOne(model.UserAuth{}, "id", id)
	if !user.IsEmpty() && user.Password == util.Generator.MD5(pwdVO.OldPassword) {
		user.Password = util.Generator.MD5(pwdVO.NewPassword)
		dao.Updates(&user, "password")
	} else {
		return r.OldPwdError
	}
	return r.SUCCESS
}

func (*UserService) SendCode(username string) int {
	code := util.Generator.ValidateCode()
	content := fmt.Sprintf(`
		<div style="text-align:center">
			<div>你好！欢迎访问我博客！</div>
			<div style="padding: 8px 40px 8px 50px;">
            	<p>
					您本次的验证码为
					<p style="font-size:75px;font-weight:blod;"> %s </p>
					为了保证账号安全，验证码有效期为15分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用~
				</p>
       	 	</div>
			<div>
            	<p>发送邮件专用邮箱，请勿回复。</p>
        	</div>
		</div>
	`, code)
	emailConf := config.EMailConf
	m := gomail.NewMessage()
	m.SetHeader(`From`, emailConf.Sender, "水平线之下")
	m.SetHeader(`To`, username)
	m.SetHeader(`Subject`, emailConf.Title)
	m.SetBody(emailConf.BodyType, content)
	err := gomail.NewDialer(emailConf.SMTPAddr, emailConf.SMTPPort, emailConf.Sender, emailConf.AuthCode).DialAndSend(m)
	if err != nil {
		return r.MailSendFailure
	}
	rdb.Set("code:"+username, code, time.Duration(emailConf.ExpireTime)*time.Minute)
	return r.SUCCESS
}

func (*UserService) SaveEmail(c *gin.Context, data model.EmailVO) int {
	if data.Code != rdb.Get("code:"+data.Email).Val() {
		return r.CodeError
	}
	uid, _ := c.Get("userInfoId")
	userDao.UpdateUsernameAndEmail(uid.(int), data.Email)
	dao.Updates(&model.UserAuth{Username: data.Email}, "username")
	return r.SUCCESS
}

func (*UserService) Register(data model.UserVO) int {
	if code := checkUser(data, true); code != r.SUCCESS {
		return code
	}
	userInfo := &model.UserInfo{
		Email:    data.Username,
		Nickname: "用户" + data.Username,
		Avatar:   blogInfoService.GetWebsiteConfig().UserAvatar,
	}
	dao.Create(&userInfo)
	dao.Create(&model.UserRole{
		UserId: userInfo.ID,
		RoleId: 2,
	})
	dao.Create(&model.UserAuth{
		UserInfoId:    userInfo.ID,
		Username:      data.Username,
		Password:      util.Generator.MD5(data.Password),
		LoginType:     1,
		LastLoginTime: time.Now(),
	})
	return r.SUCCESS
}

func (*UserService) UpdatePassword(data model.UserVO) int {
	if code := checkUser(data, false); code != r.SUCCESS {
		return code
	}
	userDao.UpdateUserAuthByUsername(data)
	return r.SUCCESS
}

func checkUser(data model.UserVO, flag bool) int { // flag: true -> 判断是否已经注册 false -> 判断 是否存在
	if data.Code != rdb.Get("code:"+data.Username).Val() {
		return r.CodeError
	}
	user := dao.GetOne(model.UserAuth{}, "username = ?", data.Username)
	if !user.IsEmpty() && flag { // 存在
		return r.EmailHasBeenRegistered
	} else if user.IsEmpty() && flag == false {
		return r.EmailHasNotBeenRegistered
	}
	return r.SUCCESS
}

func LoadUserByUsername(c *gin.Context, username string) (model.UserDetailDTO, int) {
	userAuth := dao.GetOne(model.UserAuth{}, "username = ?", username)
	if userAuth.IsEmpty() {
		return model.UserDetailDTO{}, r.UsernameNotExist
	}
	return convertUserDetail(userAuth, c), r.SUCCESS
}

func convertUserDetail(user model.UserAuth, c *gin.Context) model.UserDetailDTO {
	userInfo := dao.GetOne(model.UserInfo{}, "id", user.UserInfoId)
	roleList := roleDao.RoleListByUserInfoId(userInfo.ID)
	articleLikeSet := rdb.SMembers(ArticleUserLike + strconv.Itoa(userInfo.ID)).Val()
	commentLikeSet := rdb.SMembers(CommentUserLike + strconv.Itoa(userInfo.ID)).Val()
	talkLikeSet := rdb.SMembers(TalkUserLike + strconv.Itoa(userInfo.ID)).Val()
	ipAddress := util.IpUtil.GetIpAddress(c)
	ipSource := util.IpUtil.GetIpSourceSimpleIdle(ipAddress)
	userAgent := util.IpUtil.GetUserAgent(c)
	return model.UserDetailDTO{
		UserInfoDTO: model.UserInfoDTO{
			ID:             user.ID,
			UserInfoId:     userInfo.ID,
			Email:          userInfo.Email,
			LoginType:      user.LoginType,
			Username:       user.Username,
			Nickname:       userInfo.Nickname,
			Avatar:         userInfo.Avatar,
			Intro:          userInfo.Intro,
			WebSite:        userInfo.WebSite,
			ArticleLikeSet: articleLikeSet,
			CommentLikeSet: commentLikeSet,
			TalkLikeSet:    talkLikeSet,
			IpAddress:      ipAddress,
			IpSource:       ipSource,
			LastLoginTime:  time.Now(),
		},
		Password:  user.Password,
		RoleList:  roleList,
		IsDisable: userInfo.IsDisable,
		Browser:   userAgent.Name + " " + userAgent.Version.String(),
		Os:        userAgent.OS + " " + userAgent.OSVersion.String(),
	}
}

func updateLogin(userInfoDTO model.UserInfoDTO) {
	userAuth := model.UserAuth{
		Universal:     model.Universal{ID: userInfoDTO.ID},
		IpAddress:     userInfoDTO.IpAddress,
		IpSource:      userInfoDTO.IpSource,
		LastLoginTime: userInfoDTO.LastLoginTime,
	}
	dao.Updates(&userAuth, "ip_address", "ip_source", "last_login_time")
}
