package requesthandler

import (
	"sheddit/scheduleDatabase"
	"sheddit/taskscheduler"
	ty "sheddit/types"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	ErrorMsg string
}

func schedulePost(c *gin.Context) {
	req := ty.ScheduleRequest{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, ErrorResponse{
			"Invalid_Format",
		})
	} else {
		scheduleDatabase.AddToSchedule(&req)
		taskscheduler.SchedulePost(&req)
		c.JSON(200, "OK")
	}
}
