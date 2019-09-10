package integral

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func ListIntegral(c *gin.Context)  {
	fmt.Printf("List integral...")
}

func ShowIntegral(c *gin.Context)  {
	userId := c.Param("userId")
	fmt.Printf("Show %s integral.", userId)
}