package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"myblog/dao"
	"myblog/model"
	"myblog/util"
	"myblog/util/r"
	"strings"
	"time"
)

const (
	User   = "user:"
	Delete = "delete:"
)

var rdb = dao.GetRDB()

func ListenOnline() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, _ := c.Get("uuid")
		uuid := uid.(string)
		session := sessions.Default(c)
		var sessionInfo model.SessionInfo
		// 不空就直接从redis取，还有就是若是被强制下线了这个kv过期时间将与token过期时间长短一致，所以被强制下线时基本上都能够取到
		if str, err := rdb.Get(User + uuid).Result(); err == nil {
			util.Json.Unmarshal(str, &sessionInfo)
		} else if str, err := rdb.Get(Delete + uuid).Result(); err == nil {
			util.Json.Unmarshal(str, &sessionInfo)
		} else {
			sessionInfo = session.Get(User + uuid).(model.SessionInfo)
		}
		if strings.Index(c.FullPath(), "logout") != -1 { // 退出
			session.Delete(User + uuid)
			session.Save()
			rdb.Del(User + uuid)
			c.Abort()
			return
		}
		if sessionInfo.IsOffline == 1 { // 被强制下线
			session.Delete(User + uuid)
			session.Save()
			rdb.Del(Delete + uuid)
			r.Send(c, r.ForceOffline)
			c.Abort()
			return
		}
		c.Next()
		rdb.Set(User+uuid, util.Json.Marshal(sessionInfo), 10*time.Minute) // 更新在线状态
	}
}
