package middleware

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func InitAuthStore(r *gin.Engine) {
	serviceDisc, deployed := os.LookupEnv("COPILOT_SERVICE_DISCOVERY_ENDPOINT")
	var redisHost string
	if deployed {
		redisHost = fmt.Sprintf("redis.%s:6379", serviceDisc)
	} else {
		redisHost = "localhost:6379"
	}

	store, error := redis.NewStore(10, "tcp",
		redisHost, "",
		[]byte(os.Getenv("AUTH_CLIENT_SECRET")))
	if error != nil {
		log.Panic(error)
	}
	r.Use(sessions.Sessions("auth", store))
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		session.Set("test", "value")
		session.Save()
		c.Next()
	}
}
