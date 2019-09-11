package web

import (
	"Journey/app/integral"
	"Journey/app/tag"
	"github.com/gin-gonic/gin"
)

func InitRouter() * gin.Engine {
	r := gin.New()
	v1 := r.Group("/api/v1")

	integral.RegisterSubRouters(v1.Group("/integral"))
	tag.RegisterSubRouters(v1.Group("/tag"))

	return r
}