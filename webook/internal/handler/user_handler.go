package handler

import (
	"Project-WeBook/webook/internal/domain"
	"Project-WeBook/webook/internal/service"
	"errors"
	"net/http"
	"time"

	"github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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
		emailRegexPattern    = `^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`
		passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,72}$`
		birthdayRegexPattern = `^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$`
	)

	emailExp := regexp2.MustCompile(emailRegexPattern, 0)
	passwordExp := regexp2.MustCompile(passwordRegexPattern, 0)
	birthdayExp := regexp2.MustCompile(birthdayRegexPattern, 0)

	return &UserHandler{
		svc:         svc,
		emailExp:    emailExp,
		passwordExp: passwordExp,
		birthdayExp: birthdayExp,
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
		ctx.String(http.StatusOK, "系统错误(注册邮箱校验)")
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
		ctx.String(http.StatusOK, "系统错误(注册密码校验)")
		return
	}

	if !ok {
		ctx.String(http.StatusOK, "密码必须大于8位，包含数字、特殊字符")
		return
	}

	//调用 svc 的方法进行数据库操作
	err = user.svc.SignUp(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if errors.Is(err, service.ErrUserDuplicateEmail) {
		ctx.String(http.StatusOK, "该邮箱已注册")
		return
	}
	if err != nil {
		ctx.String(http.StatusOK, "系统错误(注册信息数据库存储)")
		return
	}

	//成功响应
	ctx.String(http.StatusOK, "注册成功")

}

// LoginJWT 使用JWT
func (user *UserHandler) LoginJWT(ctx *gin.Context) {
	type LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req LoginReq
	if err := ctx.Bind(&req); err != nil {
		return
	}

	// 校验密码
	u, err := user.svc.Login(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if errors.Is(err, service.ErrInvalidUserOrPassword) {
		ctx.String(http.StatusOK, "用户名或密码错误")
		return
	}
	if err != nil {
		ctx.String(http.StatusOK, "系统错误(密码校验)")
		return
	}

	// 登录成功
	claims := UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),
		},
		UserId:    u.Id,
		UserAgent: ctx.Request.UserAgent(),
	}
	// 设置jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenStr, err := token.SignedString([]byte("9sK7$pR2!zG5&qB8@tN3#mC6%vH1*dJ4"))
	if err != nil {
		ctx.String(http.StatusInternalServerError, "系统错误")
		return
	}
	ctx.Header("x-jwt-token", tokenStr)
	ctx.String(http.StatusOK, "登录成功")
}

// Login 登陆
func (user *UserHandler) Login(ctx *gin.Context) {
	type LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req LoginReq
	if err := ctx.Bind(&req); err != nil {
		return
	}

	// 校验密码
	u, err := user.svc.Login(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if errors.Is(err, service.ErrInvalidUserOrPassword) {
		ctx.String(http.StatusOK, "用户名或密码错误")
		return
	}
	if err != nil {
		ctx.String(http.StatusOK, "系统错误(密码校验)")
		return
	}

	// 登录成功
	// 设置 session
	session := sessions.Default(ctx)
	// 设置在 session 中的值
	session.Set("userId", u.Id)
	// cookie 安全属性
	session.Options(sessions.Options{
		// 仅 https 传输(生产环境用)
		Secure: true,
		// 仅 http 传输
		HttpOnly: true,
		MaxAge:   60,
	})
	err = session.Save()
	if err != nil {
		ctx.String(http.StatusOK, "系统错误(save session)")
		return
	}
	ctx.String(http.StatusOK, "登录成功")
}

// Logout 登出
func (user *UserHandler) Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Options(sessions.Options{
		MaxAge: -1,
	})
	err := session.Save()
	if err != nil {
		ctx.String(http.StatusOK, "系统错误(save session)")
		return
	}
	ctx.String(http.StatusOK, "已退出登录")
}

// Edit 编辑
func (user *UserHandler) Edit(ctx *gin.Context) {
	type EditReq struct {
		Name      string `json:"name"`
		Birthday  string `json:"birthday"`
		Biography string `json:"biography"`
	}
	var req EditReq
	if err := ctx.Bind(&req); err != nil {
		return
	}

	// 昵称长度校验
	if len(req.Name) >= 16 {
		ctx.String(http.StatusOK, "昵称过长(不要超过16个字节)")
		return
	}

	// 生日格式校验
	ok, err := user.birthdayExp.MatchString(req.Birthday)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误(生日校验)")
		return
	}
	if !ok {
		ctx.String(http.StatusOK, "生日格式为(xxxx-xx-xx)")
		return
	}

	// 个人简介长度校验
	if len(req.Biography) >= 80 {
		ctx.String(http.StatusOK, "简介过长(不要超过80个字节)")
		return
	}

	// 更新数据库
	//session := sessions.Default(ctx)
	//id := session.Get("userId")
	userId, ok := ctx.Get("userid")
	if !ok {
		// 正常会拿到userid
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	uid, ok := userId.(int64)
	if !ok {
		// 断言出问题
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	err = user.svc.Edit(ctx, domain.User{
		Id:        uid,
		Name:      req.Name,
		Birthday:  req.Birthday,
		Biography: req.Biography,
	})
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}

	ctx.String(http.StatusOK, "修改成功")
}

// ProFile 用户信息
func (user *UserHandler) ProFile(ctx *gin.Context) {

	// 查看 id 对应个人信息
	//session := sessions.Default(ctx)
	//id := (session.Get("userId")).(int64)

	userid, ok := ctx.Get("userId")
	if !ok {
		// 正常会拿到userid
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	uid, ok := userid.(int64)
	if !ok {
		// 断言出问题
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	u, err := user.svc.GetProfile(ctx, uid)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}

	type User struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Birth string `json:"birth"`
		Bio   string `json:"bio"`
	}
	ctx.JSON(http.StatusOK, User{
		Name:  u.Name,
		Email: u.Email,
		Birth: u.Birthday,
		Bio:   u.Biography,
	})

}

func (user *UserHandler) Test(ctx *gin.Context) {
	ctx.String(http.StatusOK, "测试成功,你已登录")
}

type UserClaims struct {
	jwt.RegisteredClaims
	// 声明要放到token的数据
	UserId    int64
	UserAgent string
}
