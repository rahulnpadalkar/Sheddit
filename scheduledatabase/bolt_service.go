package scheduledatabase

import (
	"encoding/json"
	"os"
	"reflect"
	"sheddit/logger"
	"sheddit/types"
	"sheddit/utils"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
	_ "github.com/joho/godotenv/autoload"
)

const ISOFormat = "2006-01-02T15:04:05.999999999Z07:00"

type BoltService struct {
	DB *bolt.DB
}

// InitializeDB : setup db
func (boltDB *BoltService) InitializeDB() error {
	db, err := bolt.Open("schedule.db", 0600, nil)
	if err != nil {
		logger.Log(err.Error())
		return err
	}
	logger.Log("DB initialized")
	boltDB.DB = db
	return nil
}

// AddToSchedule : Adds a new scheduled post to the database
func (boltDB *BoltService) AddToSchedule(schedulePost *types.ScheduleRequest) error {
	boltDB.DB.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(os.Getenv("bucketname")))
		if err != nil {
			logger.Log(err.Error())
			return err
		}
		scheduleID, err := b.NextSequence()
		if err != nil {
			logger.Log(err.Error())
			return err
		}
		schedulePost.ScheduleID = strconv.FormatUint(scheduleID, 10)
		schedulePost.Complete = false
		jsonstruct, err := json.Marshal(schedulePost)
		if err != nil {
			logger.Log(err.Error())
			return err
		}
		id := utils.ConvertIntToByte(int(scheduleID))
		b.Put(id, jsonstruct)
		return nil
	})
	return nil
}

// UpdateStatus : Updates status of a schedule to complete
func (boltDB *BoltService) UpdateStatus(scheduleRequest interface{}) {
	schedulePost := types.ScheduleRequest{}
	var scheduleID string
	switch reflect.TypeOf(scheduleRequest).Name() {
	case "ScheduleRequest":
		temp := scheduleRequest.(types.ScheduleRequest)
		scheduleID = temp.ScheduleID
	case "EmailRequest":
		temp := scheduleRequest.(types.EmailRequest)
		scheduleID = temp.ScheduleID
	}
	boltDB.DB.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(os.Getenv("bucketname")))
		if err != nil {
			logger.Log(err.Error())
			os.Exit(1)
		}
		jsonstruct := b.Get([]byte(scheduleID))
		json.Unmarshal(jsonstruct, &schedulePost)
		schedulePost.Complete = true
		updatedjsonstruct, err := json.Marshal(schedulePost)
		if err != nil {
			logger.Log(err.Error())
			os.Exit(1)
		}
		b.Put([]byte(scheduleID), updatedjsonstruct)
		return nil
	})
}

// AddEmailRequest : Add a new email request to the database
func (boltDB *BoltService) AddEmailRequest(email *types.EmailRequest) error {
	boltDB.DB.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(os.Getenv("emailrequests")))
		if err != nil {
			logger.Log(err.Error())
			return err
		}
		scheduleID, err := b.NextSequence()
		if err != nil {
			logger.Log(err.Error())
			return err
		}
		email.ScheduleID = strconv.FormatUint(scheduleID, 10)
		email.Complete = false
		jsonstruct, err := json.Marshal(email)
		if err != nil {
			logger.Log(err.Error())
			return err
		}
		id := utils.ConvertIntToByte(int(scheduleID))
		b.Put(id, jsonstruct)
		return nil
	})
	return nil
}

// RecoverSchedules : Schedule posts which aren't completed and whose scheudle time hasn't passed.
func (boltDB *BoltService) RecoverSchedules() []types.ScheduleRequest {
	var recoverSchedules []types.ScheduleRequest
	boltDB.DB.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(os.Getenv("bucketname")))
		if err != nil {
			logger.Log(err.Error())
			os.Exit(1)
		}
		cur := b.Cursor()

		for k, post := cur.First(); k != nil; k, post = cur.Next() {
			scheduledPost := types.ScheduleRequest{}
			json.Unmarshal(post, &scheduledPost)
			scheduleTime, err := time.Parse(ISOFormat, scheduledPost.ScheduleDate)
			if err != nil {
				logger.Log(err.Error())
				os.Exit(1)
			}
			if !scheduledPost.Complete && !time.Now().After(scheduleTime) {
				recoverSchedules = append(recoverSchedules, scheduledPost)
			}
		}
		return nil
	})
	return recoverSchedules
}

// GetAllSchedules : This returns all schedules
func (boltDB *BoltService) GetAllSchedules() []types.ScheduleRequest {
	var scheduledPosts []types.ScheduleRequest
	boltDB.DB.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(os.Getenv("bucketname")))
		if err != nil {
			logger.Log(err.Error())
		}
		cur := b.Cursor()

		for k, post := cur.First(); k != nil; k, post = cur.Next() {
			scheduledPost := types.ScheduleRequest{}
			json.Unmarshal(post, &scheduledPost)
			scheduledPosts = append(scheduledPosts, scheduledPost)
		}
		return nil
	})
	return scheduledPosts
}

// GetCompleteSchedules : Returns all complete schedules
func (boltDB *BoltService) GetCompleteSchedules() []types.ScheduleRequest {

	allSchedules := boltDB.GetAllSchedules()
	var completeSchedules []types.ScheduleRequest
	for _, post := range allSchedules {
		if post.Complete {
			completeSchedules = append(completeSchedules, post)
		}
	}
	return completeSchedules
}

//GetIncompleteSchedules : Returns all incomplete schedules
func (boltDB *BoltService) GetIncompleteSchedules() []types.ScheduleRequest {

	allSchedules := boltDB.GetAllSchedules()
	var incompleteSchedules []types.ScheduleRequest
	for _, post := range allSchedules {
		if !post.Complete {
			incompleteSchedules = append(incompleteSchedules, post)
		}
	}
	return incompleteSchedules
}

// DropTables : Delete buckets in the DB
func (boltDB *BoltService) DropTables() {
	buckets := []string{os.Getenv("bucketname")}
	boltDB.DB.Update(func(tx *bolt.Tx) error {
		for _, tablename := range buckets {
			err := tx.DeleteBucket([]byte(tablename))
			if err != nil {
				logger.Log("Error Deleting Bucket :" + err.Error())
				err = nil
			}
		}
		return nil
	})
}

// RecoverEmailSchedules : Recover email schedules after the server restarts from a crash
func (boltDB *BoltService) RecoverEmailSchedules() []types.EmailRequest {
	var recoverSchedules []types.EmailRequest
	boltDB.DB.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(os.Getenv("emailrequests")))
		if err != nil {
			logger.Log(err.Error())
			os.Exit(1)
		}
		cur := b.Cursor()

		for k, post := cur.First(); k != nil; k, post = cur.Next() {
			emailRequest := types.EmailRequest{}
			json.Unmarshal(post, &emailRequest)
			scheduleTime, err := time.Parse(ISOFormat, emailRequest.ScheduleDate)
			if err != nil {
				logger.Log(err.Error())
				os.Exit(1)
			}
			if !emailRequest.Complete && !time.Now().After(scheduleTime) {
				recoverSchedules = append(recoverSchedules, emailRequest)
			}
		}
		return nil
	})
	return recoverSchedules
}

// GetAllEmailSchedules : Get all email requests
func (boltDB *BoltService) GetAllEmailSchedules() []types.EmailRequest {
	var scheduledPosts []types.EmailRequest
	boltDB.DB.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(os.Getenv("email_bucket")))
		if err != nil {
			logger.Log(err.Error())
		}
		cur := b.Cursor()

		for k, post := cur.First(); k != nil; k, post = cur.Next() {
			scheduledPost := types.EmailRequest{}
			json.Unmarshal(post, &scheduledPost)
			scheduledPosts = append(scheduledPosts, scheduledPost)
		}
		return nil
	})
	return scheduledPosts
}
