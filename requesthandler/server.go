package requesthandler

import (
	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()
	addcustomvalidator()
	router.GET("/status", status)
	router.POST("/email", scheduleEmail)
	router.POST("/schedulePost", schedulePost)
	router.GET("/getallschedules", getAllSchedules)
	router.Run(":7009")
}
