package integral

import "github.com/gin-gonic/gin"

func RegisterSubRouters(router *gin.RouterGroup)  {
	router.GET("/$", ListIntegral)
	router.GET("/:userId", ShowIntegral)
}
