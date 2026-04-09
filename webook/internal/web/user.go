package web

import (
	"net/http"

	"github.com/dlclark/regexp2"

	"github.com/gin-gonic/gin"
)

// UserHandler 用户处理程序,在他上面定义和用户有关系的路由
type UserHandler struct {
	emailExp    *regexp2.Regexp
	passwordExp *regexp2.Regexp
}

func NewUserHandler() *UserHandler {
	// 正则表达式,用于校验
	const (
		emailRegexPattern    = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
		passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
	)

	emailExp := regexp2.MustCompile(emailRegexPattern, 0)
	passwordExp := regexp2.MustCompile(passwordRegexPattern, 0)

	return &UserHandler{
		emailExp:    emailExp,
		passwordExp: passwordExp,
	}
}

// SignUp 注册
func (u *UserHandler) SignUp(ctx *gin.Context) {
	type SignUpReq struct {
		Email           string `json:"email"`
		ConfirmPassword string `json:"confirmPassword"`
		Password        string `json:"password"`
	}

	//实例化
	var req SignUpReq
	// Bind 方法会根据 Content-Type 来自动解析你的前端数据到 req 里面
	// 解析错了，就会直接写回一个 400 的错误
	if err := ctx.Bind(&req); err != nil {
		return
	}

	// 邮箱校验
	ok, err := u.emailExp.MatchString(req.Email)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}

	if !ok {
		ctx.String(http.StatusOK, "你的邮箱格式不对")
		return
	}

	// 密码校验
	if req.ConfirmPassword != req.Password {
		ctx.String(http.StatusOK, "两次输入的密码不一致")
		return
	}

	ok, err = u.passwordExp.MatchString(req.Password)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}

	if !ok {
		ctx.String(http.StatusOK, "密码必须大于8位，包含数字、特殊字符")
		return
	}

	//成功响应
	ctx.String(http.StatusOK, "注册成功")

}

// Login 登陆
func (u *UserHandler) Login(ctx *gin.Context) {

}

// Edit 编辑
func (u *UserHandler) Edit(ctx *gin.Context) {

}

// ProFile 用户信息
func (u *UserHandler) ProFile(ctx *gin.Context) {

}
