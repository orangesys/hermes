package routers

import (
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	v1 "github.com/orangesys/hermes/routers/api/v1"
	uuid "github.com/satori/go.uuid"
)

func RequestIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Request-Id", uuid.NewV4().String())
		c.Next()
	}
}

func RevisionMiddleware() gin.HandlerFunc {
	revision := os.Getenv("REVISION")
	if revision == "" {
		log.Println("can not get revision from env")
		return func(c *gin.Context) {
			c.Next()
		}
	}
	revision = strings.TrimSpace(string(revision))
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Revision", revision)
		c.Next()
	}
}

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(RequestIdMiddleware())
	r.Use(RevisionMiddleware())
	r.HEAD("/ping", v1.Ping)

	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/user", v1.CreateUser)
	}
	return r
}
