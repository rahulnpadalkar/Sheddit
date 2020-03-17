package main

import (
	logger "sheddit/logger"
	server "sheddit/requesthandler"
	schedulerdb "sheddit/scheduledatabase"
	taskscheduler "sheddit/taskscheduler"
)

func main() {
	logger.InitialzeLogger()
	schedulerdb.InitializeDB()
	recoverSchedues := schedulerdb.RecoverSchedules()
	for _, schedule := range recoverSchedues {
		taskscheduler.SchedulePost(&schedule)
	}
	server.StartServer()
	defer logger.CloseFile()
}
