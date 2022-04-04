package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Auth(r *gin.Engine) gin.HandlerFunc {
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
		Secure:   true,
		HttpOnly: true,
	})
	r.Use(sessions.Sessions("auth", store))

	return func(c *gin.Context) {
		session := sessions.Default(c)
		session.Set("test", "value")
		session.Save()
		c.Next()
	}
}
