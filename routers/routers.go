package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/orangesys/hermes/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/user", v1.CreateUser)
	}
	return r
}
