package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @program:   go-basic-exercises
// @file:      login.go
// @author:    bowen
// @time:      2023-09-23 11:32
// @description:

type LoginMiddleWareBuilder struct {
	paths []string
}

func NewLoginMiddlerwareBuilder() *LoginMiddleWareBuilder {
	return &LoginMiddleWareBuilder{}
}

func (l *LoginMiddleWareBuilder) IgnorePaths(path string) *LoginMiddleWareBuilder {
	l.paths = append(l.paths, path)
	return l

}
func (l *LoginMiddleWareBuilder) Buil() gin.HandlerFunc {

	return func(c *gin.Context) {
		for _, path := range l.paths {
			if c.Request.URL.Path == path {
				return
			}
		}

		sess := sessions.Default(c)
		if sess == nil {
			// 没有登录
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		id := sess.Get("userid")
		if id == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
