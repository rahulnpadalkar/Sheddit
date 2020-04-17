package scheduledatabase

import (
	"os"
	"sheddit/logger"
	"sheddit/types"
	"time"

	"reflect"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/twinj/uuid"
)

// PostgresClinet : Postgres client struct
type PostgresClient struct {
	DB *pg.DB
}

var models = []interface{}{(*types.ScheduleRequest)(nil), (*types.EmailRequest)(nil)}

var pgclient *PostgresClient

// InitializeDB : Initiliaze PostgresSQL database
func (pgclient *PostgresClient) InitializeDB() error {
	options, err := pg.ParseURL(os.Getenv("postgres_url"))
	if err != nil {
		logger.Log(err.Error())
		return err
	}
	pgclient.DB = pg.Connect(options)

	//Create tables corresponding to request types

	for _, model := range models {
		pgclient.DB.CreateTable(model, &orm.CreateTableOptions{
			IfNotExists: true,
		})
	}
	return nil
}

// AddToSchedule : Add schedule to database
func (pgclient *PostgresClient) AddToSchedule(schedule *types.ScheduleRequest) error {
	schedule.Complete = false
	schedule.ScheduleID = uuid.NewV4().String()
	err := pgclient.DB.Insert(schedule)
	if err != nil {
		logger.Log(err.Error())
		return err
	}
	return nil
}

// AddEmailRequest : Add email request to database
func (pgclient *PostgresClient) AddEmailRequest(email *types.EmailRequest) error {
	email.Complete = false
	email.ScheduleID = uuid.NewV4().String()
	err := pgclient.DB.Insert(email)

	if err != nil {
		logger.Log(err.Error())
		return err
	}
	return nil
}

// UpdateStatus : Updates status of a scheduled post
func (pgclient *PostgresClient) UpdateStatus(scheduleRequest interface{}) {
	switch reflect.TypeOf(scheduleRequest).Name() {
	case "ScheduleRequest":
		temp := scheduleRequest.(types.ScheduleRequest)
		temp.Complete = true
		if err := pgclient.DB.Update(&temp); err != nil {
			logger.Log("Couldn't update status beacuse: " + err.Error())
		}
	case "EmailRequest":
		temp := scheduleRequest.(types.EmailRequest)
		temp.Complete = true
		if err := pgclient.DB.Update(&temp); err != nil {
			logger.Log("Couldn't update status beacuse: " + err.Error())
		}
	}
}

// RecoverSchedules : Recover ScheduleRequest after the server restarts after a crash
func (pgclient *PostgresClient) RecoverSchedules() []types.ScheduleRequest {
	var requests, filteredRequest []types.ScheduleRequest
	if _, err := pgclient.DB.Query(&requests, `SELECT * FROM schedule_requests WHERE complete = false`); err != nil {
		logger.Log("Recovering Scheudle FAILED : " + err.Error())
	}
	for _, schedule := range requests {
		if scheduleTime, _ := time.Parse(ISOFormat, schedule.ScheduleDate); !time.Now().After(scheduleTime) {
			filteredRequest = append(filteredRequest, schedule)
		}
	}
	return filteredRequest
}

// GetAllSchedules : Get All Schedules
func (pgclient *PostgresClient) GetAllSchedules() []types.ScheduleRequest {
	var requests []types.ScheduleRequest
	if _, err := pgclient.DB.Query(&requests, `SELECT * FROM schedule_requests`); err != nil {
		logger.Log("Recovering Scheudle FAILED : " + err.Error())
	}
	return requests
}

// GetCompleteSchedules : Get all completed schedules
func (pgclient *PostgresClient) GetCompleteSchedules() []types.ScheduleRequest {
	var requests []types.ScheduleRequest
	if _, err := pgclient.DB.Query(&requests, `SELECT * FROM schedule_requests WHERE complete = true`); err != nil {
		logger.Log("Recovering Scheudle FAILED : " + err.Error())
	}
	return requests
}

// GetIncompleteSchedules : Get all incomplete schedules
func (pgclient *PostgresClient) GetIncompleteSchedules() []types.ScheduleRequest {
	var requests []types.ScheduleRequest
	if _, err := pgclient.DB.Query(&requests, `SELECT * FROM schedule_requests WHERE complete = false`); err != nil {
		logger.Log("Recovering Scheudle FAILED : " + err.Error())
	}
	return requests
}

// DropTables : Drop table from database
func (pgclient *PostgresClient) DropTables() {
	tables := []string{"email_requests", "schedule_requests"}
	for _, tableName := range tables {
		query := "DROP TABLE IF EXISTS " + tableName
		if _, err := pgclient.DB.Query(nil, query); err != nil {
			logger.Log("Error droping table : " + err.Error())
		}
	}
}

// RecoverEmailSchedules : Recover email schedules after the server restarts from a crash
func (pgclient *PostgresClient) RecoverEmailSchedules() []types.EmailRequest {
	var requests, filteredRequest []types.EmailRequest
	if _, err := pgclient.DB.Query(&requests, `SELECT * FROM email_requests WHERE complete = false`); err != nil {
		logger.Log("Recovering Scheudle FAILED : " + err.Error())
	}
	for _, schedule := range requests {
		if scheduleTime, _ := time.Parse(ISOFormat, schedule.ScheduleDate); !time.Now().After(scheduleTime) {
			filteredRequest = append(filteredRequest, schedule)
		}
	}
	return filteredRequest
}

// GetAllEmailSchedules : Get All Schedules
func (pgclient *PostgresClient) GetAllEmailSchedules() []types.EmailRequest {
	var requests []types.EmailRequest
	if _, err := pgclient.DB.Query(&requests, `SELECT * FROM email_requests`); err != nil {
		logger.Log("Recovering Scheudle FAILED : " + err.Error())
	}
	return requests
}
