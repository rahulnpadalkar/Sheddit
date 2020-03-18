package taskscheduler

import (
	"log"
	"sheddit/actions"
	scheduledb "sheddit/scheduledatabase"
	"sheddit/types"
	"time"
)

const ISOFormat = "2006-01-02T15:04:05.999999999Z07:00"

// SchedulePost : Schedule a reddit post
func SchedulePost(schedulePost *types.ScheduleRequest) {
	scheduleTime, err := time.Parse(ISOFormat, schedulePost.ScheduleDate)
	if err != nil {
		log.Fatal(err)
	}
	timeDuration := time.Until(scheduleTime)
	time.AfterFunc(timeDuration, func() {
		success, _ := actions.BulkPost(*schedulePost)
		if success {
			scheduledb.UpdateStatus(schedulePost.ScheduleID)
		}
	})
}
