package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"go-basic-exercises/webook/internal/repository"
	"go-basic-exercises/webook/internal/repository/dao"
	"go-basic-exercises/webook/internal/service"
	"go-basic-exercises/webook/internal/web"
	"go-basic-exercises/webook/internal/web/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

// @program:   go-basic-exercises
// @file:      main.go
// @author:    bowen
// @time:      2023-09-19 14:58
// @description:

func main() {
	db := initDB()
	server := initWebServer()
	u := initUser(db)
	u.RegisterRoutes(server)
	//u.RegisterRoutesV1(server.Group("/user"))

	server.Run(":8080")
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:815qza@tcp(192.168.72.130)/webook"))
	if err != nil {
		// 只会在初始化过程panic，panic相当于整个goroutine结束，一旦初始化过程出错，就不要启动了。
		panic(err)
	}

	err = dao.InitTable(db)
	if err != nil {
		panic(err)
	}
	return db

}

func initWebServer() *gin.Engine {
	server := gin.Default()
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				return true
			}
			return strings.Contains(origin, "bowen.love")
		},
		MaxAge: 12 * time.Hour,
	}))

	store := cookie.NewStore([]byte("secret"))
	server.Use(sessions.Sessions("mysession", store))
	//这种写法要记得
	server.Use(middleware.NewLoginMiddlerwareBuilder().IgnorePaths("/users/signup").IgnorePaths("/users/login").Buil())
	return server

}

func initUser(db *gorm.DB) *web.UserHandler {
	ud := dao.NewUserDao(db)
	repo := repository.NewUserRepostitory(ud)
	svc := service.NewUserService(repo)
	u := web.NewUserHandler(svc)
	return u
	//u.RegisterRout es(server)

}
