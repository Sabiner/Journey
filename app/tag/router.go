package tag

import "github.com/gin-gonic/gin"

func RegisterSubRouters(router *gin.RouterGroup)  {
	router.GET("/$", ListTag)
}
