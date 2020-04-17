package scheduledatabase

import (
	"sheddit/types"
)

// DBService : Interface which defines DBService 2 implementations Postgres an BoltDB service
type DBService interface {
	InitializeDB() error
	AddToSchedule(scheuleRequest *types.ScheduleRequest) error
	AddEmailRequest(email *types.EmailRequest) error
	UpdateStatus(scheduleRequest interface{})
	RecoverSchedules() []types.ScheduleRequest
	GetAllSchedules() []types.ScheduleRequest
	GetCompleteSchedules() []types.ScheduleRequest
	GetIncompleteSchedules() []types.ScheduleRequest
	DropTables()
	RecoverEmailSchedules() []types.EmailRequest
	GetAllEmailSchedules() []types.EmailRequest
}
