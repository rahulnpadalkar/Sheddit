package requesthandler

import (
	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()
	router.POST("/schedulePost", schedulePost)
	router.Run(":7009")
}
