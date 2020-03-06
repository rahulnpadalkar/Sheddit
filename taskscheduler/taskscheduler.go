package taskscheduler

import (
	"log"
	"sheddit/actions"
	"sheddit/types"
	"strings"
	"time"
)

const ISOFormat = "2006-01-02T15:04:05.999999999Z07:00"

// SchedulePost: Schedule a reddit post
func SchedulePost(schedulePost *types.ScheduleRequest) {
	scheduleTime, err := time.Parse(ISOFormat, schedulePost.ScheduleDate)
	if err != nil {
		log.Fatal(err)
	}
	timeDuration := time.Until(scheduleTime)
	time.AfterFunc(timeDuration, func() {
		actions.BulkPost(strings.Split(schedulePost.Subreddits, ","), schedulePost.Link, schedulePost.Title, schedulePost.ScheduleID)
	})
}
