package middleware

import (
	"encoding/gob"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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
	gob.Register(time.Now())
	return func(c *gin.Context) {
		for _, path := range l.paths {
			if c.Request.URL.Path == path {
				return
			}
		}

		sess := sessions.Default(c)
		if sess == nil {
			// 没有登录 401
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		id := sess.Get("userid")
		if id == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		updateTime := sess.Get("update_time")
		sess.Set("userId", id)
		sess.Options(sessions.Options{MaxAge: 60})
		//now := time.Now().UnixMilli()
		now := time.Now()
		// 说明还没刷新过，刚登陆。
		if updateTime == nil {
			sess.Set("update_time", now)
			err := sess.Save()
			if err != nil {
				panic(err)
			}
			return
		}

		//updateTime是有的
		//updateTimeVal, ok := updateTime.(int64)
		updateTimeVal, ok := updateTime.(time.Time)
		if !ok {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		//if now-updateTimeVal > 60*1000 {
		//	sess.Set("update_time", now)
		//	sess.Save()
		//	return
		//}

		if now.Sub(updateTimeVal) > time.Second*10 {
			sess.Set("update_time", now)
			err := sess.Save()
			if err != nil {
				panic(err)
			}

			//return
		}
	}
}
