package handler

import (
	"Project-WeBook/webook/internal/domain"
	"Project-WeBook/webook/internal/service"
	"net/http"

	"github.com/dlclark/regexp2"

	"github.com/gin-gonic/gin"
)

// UserHandler 用户处理程序,在他上面定义和用户有关系的路由
type UserHandler struct {
	svc         *service.UserService
	emailExp    *regexp2.Regexp
	passwordExp *regexp2.Regexp
	birthdayExp *regexp2.Regexp
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	// 正则表达式,用于校验
	const (
		emailRegexPattern    = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
		passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}`
	)

	emailExp := regexp2.MustCompile(emailRegexPattern, 0)
	passwordExp := regexp2.MustCompile(passwordRegexPattern, 0)

	return &UserHandler{
		svc:         svc,
		emailExp:    emailExp,
		passwordExp: passwordExp,
	}
}

// SignUp 注册
func (user *UserHandler) SignUp(ctx *gin.Context) {
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
	ok, err := user.emailExp.MatchString(req.Email)
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

	ok, err = user.passwordExp.MatchString(req.Password)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}

	if !ok {
		ctx.String(http.StatusOK, "密码必须大于8位，包含数字、特殊字符")
		return
	}

	//调用 svc 的方法
	err = user.svc.SignUp(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}

	//成功响应
	ctx.String(http.StatusOK, "注册成功")

}

// Login 登陆
func (user *UserHandler) Login(ctx *gin.Context) {

}

// Edit 编辑
func (user *UserHandler) Edit(ctx *gin.Context) {

}

// ProFile 用户信息
func (user *UserHandler) ProFile(ctx *gin.Context) {

}
