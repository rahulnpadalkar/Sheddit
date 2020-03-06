package scheduleDatabase

import (
	"encoding/json"
	"os"
	"sheddit/logger"
	"sheddit/types"
	"sheddit/utils"
	"time"

	"github.com/boltdb/bolt"
	_ "github.com/joho/godotenv/autoload"
)

const ISOFormat = "2006-01-02T15:04:05.999999999Z07:00"

var db *bolt.DB

// InitializeDB : setup db
func InitializeDB() {
	var err error
	db, err = bolt.Open("schedule.db", 0600, nil)
	if err != nil {
		logger.Log(err.Error())
		os.Exit(1)
	}
	logger.Log("DB initialized")
}

// AddToSchedule : Adds a new scheduled post to the database
func AddToSchedule(schedulePost *types.ScheduleRequest) int {
	var sid int
	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(os.Getenv("bucketname")))
		if err != nil {
			logger.Log(err.Error())
			os.Exit(1)
		}
		scheduleID, err := b.NextSequence()
		if err != nil {
			logger.Log(err.Error())
			os.Exit(1)
		}
		schedulePost.ScheduleID = int(scheduleID)
		sid = int(scheduleID)
		schedulePost.Complete = false
		jsonstruct, err := json.Marshal(schedulePost)
		if err != nil {
			logger.Log(err.Error())
			os.Exit(1)
		}
		id := utils.ConvertIntToByte(int(scheduleID))
		b.Put(id, jsonstruct)
		return nil
	})
	return sid
}

// UpdateStatus : Updates status of a schedule to complete
func UpdateStatus(scheduleID int) {
	schedulePost := types.ScheduleRequest{}
	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(os.Getenv("bucketname")))
		if err != nil {
			logger.Log(err.Error())
			os.Exit(1)
		}
		scheduleIden := utils.ConvertIntToByte(scheduleID)
		jsonstruct := b.Get(scheduleIden)
		json.Unmarshal(jsonstruct, &schedulePost)
		schedulePost.Complete = true
		updatedjsonstruct, err := json.Marshal(schedulePost)
		if err != nil {
			logger.Log(err.Error())
			os.Exit(1)
		}
		b.Put(scheduleIden, updatedjsonstruct)
		return nil
	})
}

// RecoverSchedules : Schedule posts which aren't completed and whose scheudle time hasn't passed.
func RecoverSchedules() []types.ScheduleRequest {
	var recoverSchedules []types.ScheduleRequest
	db.Update(func(tx *bolt.Tx) error {
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
