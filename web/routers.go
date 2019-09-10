package web

import (
	"Journey/app/tag"
	"Journey/app/user"
	"github.com/gin-gonic/gin"
)

func InitRouter() * gin.Engine {
	r := gin.New()
	v1 := r.Group("/api/v1")

	user.RegisterSubRouters(v1.Group("/user"))
	tag.RegisterSubRouters(v1.Group("/tag"))

	return r
}