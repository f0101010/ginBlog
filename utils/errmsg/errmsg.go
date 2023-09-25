package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	// code =1000... 用户模块的错误

	ErrorUsernameUsed  = 1001
	ErrorPasswordWrong = 1002
	ErrorUserNotExist  = 1003
	ErrorTokenExist    = 1004
	ErrorTokenRuntime  = 1005
	ErrorTokenWrong    = 1006
	ErrorTokenType     = 1007
	ErrorUserNoRight   = 1008

	// code =2000... 文章模块的错误

	ErrorArtNotExist = 2002

	// code =3000... 分类模块的错误

	ErrorCatenameUsed = 3001
	ErrorCateNotExist = 3002
)

var codeMsg = map[int]string{
	SUCCESS:            "OK",
	ERROR:              "FAIL",
	ErrorUsernameUsed:  "用户名已存在",
	ErrorPasswordWrong: "密码错误",
	ErrorUserNotExist:  "用户不存在",
	ErrorTokenExist:    "TOKEN不存在",
	ErrorTokenRuntime:  "TOKEN已过期",
	ErrorTokenWrong:    "TOKEN不正确",
	ErrorTokenType:     "TOKEN格式错误",
	ErrorCatenameUsed:  "该分类已存在",
	ErrorArtNotExist:   "该文章不存在",
	ErrorCateNotExist:  "该分类不存在",
	ErrorUserNoRight:   "该用户无权限",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
