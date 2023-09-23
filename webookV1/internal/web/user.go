package web

import (
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-basic-exercises/webook/internal/domain"
	"go-basic-exercises/webook/internal/service"
	"net/http"
)

// @program:   go-basic-exercises
// @file:      user.go
// @author:    bowen
// @time:      2023-09-19 14:58
// @description:

type UserHandler struct {
	svc         *service.UserService
	emailExp    *regexp.Regexp
	passwordExp *regexp.Regexp
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	const (
		emailRegexPattern = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
		// 和上面比起来，用 ` 看起来就比较清爽
		passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{6,}$`
	)
	emailexp := regexp.MustCompile(emailRegexPattern, regexp.None)
	passexp := regexp.MustCompile(passwordRegexPattern, regexp.None)
	return &UserHandler{emailExp: emailexp, passwordExp: passexp, svc: svc}
}

func (u *UserHandler) RegisterRoutesV1(ug *gin.RouterGroup) {
	ug.GET("/profile")
	ug.POST("/signup")
	ug.POST("/edit")
	ug.POST("/login")
}

// 路由注册
func (u *UserHandler) RegisterRoutes(server *gin.Engine) {
	ug := server.Group("/users")
	ug.GET("/profile", u.Profile)
	ug.POST("/signup", u.SignUp)
	ug.POST("/edit")
	ug.POST("/login", u.Login)
	//server.POST("/user/signup", u.SignUp)
	//server.POST("/user/login", u.Login)
	//server.POST("/user/edit", u.Edit)
	//server.GET("/user/profile", u.Profile)
}

func (u *UserHandler) SignUp(c *gin.Context) {
	type SignUpReq struct {
		Email           string `json:"email"`
		ConfirmPassword string `json:"confirmPassword"`
		Password        string `json:"password"`
	}
	//fmt.Println(123)
	var req SignUpReq

	// Bind 方法根据Content Type 解析你的数据
	// 解析错了，会直接协会一个400的错误
	if err := c.Bind(&req); err != nil {
		return
	}

	//邮箱格式校验
	ok, err := u.emailExp.MatchString(req.Email)
	if err != nil {
		c.String(http.StatusOK, "系统错误")
		return
	}
	if !ok {
		c.String(http.StatusOK, "你的格式不对")
		return
	}

	// 密码和确认密码匹配
	if req.Password != req.ConfirmPassword {
		c.String(http.StatusOK, "两次输入的密码不一致")
		return
	}

	ok, err = u.emailExp.MatchString(req.Password)
	if err != nil {
		c.String(http.StatusOK, "系统错误")
		return
	}
	//密码长度校验
	if !ok {
		c.String(http.StatusOK, "密码必须大于8位，包含数字，特殊字符")
		return
	}

	err = u.svc.SignUp(c, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})

	if err == service.ErrEmailDuplicate {
		c.String(http.StatusOK, "邮箱冲突了")
	}

	if err != nil {
		c.String(http.StatusOK, "系统异常 ")
		return
	}
	c.String(http.StatusOK, "注册成功")
	//fmt.Printf("%v", req)
}

func (u *UserHandler) Login(c *gin.Context) {
	type LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// 一样的
	var req LoginReq
	if err := c.Bind(&req); err != nil {
		return
	}

	// service

	user, err := u.svc.Login(c, req.Email, req.Password)
	if err == service.InvalidUserOrPasswd {
		c.String(http.StatusOK, "用户名或密码错误")
		return
	}

	if err != nil {
		c.String(http.StatusOK, "系统错误")
		return
	}
	sess := sessions.Default(c)

	sess.Set("userid", user.ID)
	sess.Save()
	c.String(http.StatusOK, "登录成功")
	return

}

func (u *UserHandler) Edit(c *gin.Context) {

}

func (u *UserHandler) Profile(c *gin.Context) {
	c.String(http.StatusOK, "这是你的profile")
}
