package web

import (
	"Journey/app/integral"
	"Journey/app/tag"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() * gin.Engine {
	r := gin.New()
	v1 := r.Group("/api/v1")

	r.LoadHTMLGlob("frontend/**/*")
	r.Static("/semantic", "semantic")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Journey",
		})
	})

	integral.RegisterSubRouters(v1.Group("/integral"))
	tag.RegisterSubRouters(v1.Group("/tag"))

	return r
}