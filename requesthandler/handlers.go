package requesthandler

import (
	scheduleDatabase "sheddit/scheduledatabase"
	"sheddit/taskscheduler"
	ty "sheddit/types"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	ErrorMsg string
}

func status(c *gin.Context) {
	c.JSON(200, "OK")
}

func schedulePost(c *gin.Context) {
	dbInstance := scheduleDatabase.GetInstance()
	req := ty.ScheduleRequest{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, ErrorResponse{
			"Invalid_Format",
		})
	} else {
		dbInstance.AddToSchedule(&req)
		taskscheduler.SchedulePost(&req)
		c.JSON(200, "OK")
	}
}

func scheduleEmail(c *gin.Context) {
	dbInstance := scheduleDatabase.GetInstance()
	emailRequest := ty.EmailRequest{}
	if err := c.ShouldBind(&emailRequest); err != nil {
		c.JSON(400, ErrorResponse{
			"Invalid_Format",
		})
	} else {
		dbInstance.AddEmailRequest(&emailRequest)
		taskscheduler.ScheduleEmail(&emailRequest)
		c.JSON(200, "OK")
	}
}

func getAllSchedules(c *gin.Context) {
	dbInstance := scheduleDatabase.GetInstance()
	allschedules := dbInstance.GetAllSchedules()
	c.JSON(200, allschedules)
}
