package r

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SUCCESS                   = 20000
	NoLogin                   = 40001
	PasswordWrong             = 40002
	UsernameNoExist           = 40003
	Authorized                = 40300
	PermissionDenied          = 40400
	TokenCreateFail           = 40401
	TokenWrong                = 40402
	TokenFormatterError       = 40403
	TokenRuntime              = 40404
	SystemError               = 50000
	FAILURE                   = 51000
	ValidError                = 52000
	UsernameExist             = 52001
	UsernameNotExist          = 52002
	CategoryExist             = 60001
	CategoryArticleExist      = 60002
	TagExist                  = 60003
	TagArticleExist           = 60004
	OldPwdError               = 60005
	RoleExist                 = 60006
	AlbumExist                = 60007
	ForceOffline              = 60008
	TalkNotExist              = 60009
	AlbumNotExist             = 60010
	MailSendFailure           = 60011
	CodeError                 = 60012
	EmailHasBeenRegistered    = 60013
	EmailHasNotBeenRegistered = 60014
)

var codeMsg = map[int]string{
	SUCCESS:                   "操作成功",
	NoLogin:                   "用户未登录",
	PasswordWrong:             "用户名或密码错误",
	UsernameNoExist:           "用户名不存在",
	Authorized:                "没有操作权限",
	PermissionDenied:          "权限不足",
	SystemError:               "系统异常",
	FAILURE:                   "操作失败",
	ValidError:                "参数格式不正确",
	UsernameExist:             "用户名已存在",
	UsernameNotExist:          "用户名不存在",
	TokenCreateFail:           "token创建失败",
	TokenWrong:                "token错误",
	TokenFormatterError:       "token格式错误",
	TokenRuntime:              "toke已过期,请退出后重新登陆",
	CategoryExist:             "操作失败，分类名已存在",
	CategoryArticleExist:      "删除失败，该分类下存在文章",
	TagExist:                  "操作失败，标签名已存在",
	TagArticleExist:           "删除失败，该标签下存在文章",
	OldPwdError:               "旧密码不正确",
	RoleExist:                 "角色已存在",
	AlbumExist:                "相册名已存在",
	ForceOffline:              "您已被强制下线",
	TalkNotExist:              "说说不存在",
	AlbumNotExist:             "相册不存在",
	MailSendFailure:           "邮件发送失败",
	CodeError:                 "验证码错误",
	EmailHasBeenRegistered:    "邮箱已被注册",
	EmailHasNotBeenRegistered: "邮箱没有注册",
}

func getMsgByCode(code int) string {
	return codeMsg[code]
}

func Send(c *gin.Context, code int, data ...any) {
	if code != SUCCESS {
		if len(data) > 0 {
			c.JSON(http.StatusOK, gin.H{
				"flag":    false,
				"code":    code,
				"data":    data[0],
				"message": getMsgByCode(code),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"flag":    false,
			"code":    code,
			"message": getMsgByCode(code),
		})
		return
	}
	if len(data) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"flag":    true,
			"code":    code,
			"data":    data[0],
			"message": getMsgByCode(code),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"flag":    true,
		"code":    code,
		"message": getMsgByCode(code),
	})
}

func Success(c *gin.Context, data any, message ...string) {
	if len(message) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"flag":    true,
			"code":    SUCCESS,
			"data":    data,
			"message": message[0],
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"flag":    true,
		"code":    SUCCESS,
		"data":    data,
		"message": getMsgByCode(SUCCESS),
	})
}

func Failure(c *gin.Context, data any, code ...int) {
	if len(code) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"flag":    false,
			"code":    code[0],
			"data":    data,
			"message": getMsgByCode(code[0]),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"flag":    false,
		"code":    FAILURE,
		"data":    data,
		"message": getMsgByCode(FAILURE),
	})
}
