package scheduledatabase

import (
	"fmt"
	"os"
	"sheddit/types"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	InitilaizeService()
	returnCode := m.Run()
	GetInstance().DropTables()
	os.Exit(returnCode)
}
func TestAddToSchedule(t *testing.T) {
	dbInstance := GetInstance()
	dur1, _ := time.ParseDuration("180s")
	dur2, _ := time.ParseDuration("180s")
	dur3, _ := time.ParseDuration("-30s")
	postTable := []types.ScheduleRequest{
		{
			Subreddits:   "test,sandboxtest",
			Title:        "Test Post",
			Link:         "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
			ScheduleDate: (time.Now().Add(dur1).Format("2006-01-02T15:04:05.999999999Z07:00")),
			Complete:     false,
			Provider:     "Reddit",
		},
		{
			Subreddits:   "test,sandboxtest",
			Title:        "Test Post",
			Link:         "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
			ScheduleDate: (time.Now().Add(dur2).Format("2006-01-02T15:04:05.999999999Z07:00")),
			Complete:     false,
			Provider:     "Reddit",
		},
		{
			Subreddits:   "test,sandboxtest",
			Title:        "Test Post",
			Link:         "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
			ScheduleDate: (time.Now().Add(dur3).Format("2006-01-02T15:04:05.999999999Z07:00")),
			Complete:     true,
			Provider:     "Reddit",
		},
	}
	for _, post := range postTable {
		fmt.Println(dbInstance)
		dbInstance.AddToSchedule(&post)
	}

	if len(dbInstance.GetAllSchedules()) != 3 {
		t.Errorf("Expected 2 scheduled posts found  %v", len(dbInstance.GetAllSchedules()))
	}
}

func TestAddEmailRequest(t *testing.T) {
	dbInstance := GetInstance()
	dur1, _ := time.ParseDuration("180s")
	dur2, _ := time.ParseDuration("180s")
	dur3, _ := time.ParseDuration("-30s")
	postTable := []types.EmailRequest{
		{
			To:           "dummyemail@dummy.dummy",
			Template:     "hello {{name}}",
			Data:         `{name:"test"}`,
			Subject:      "Test",
			ScheduleDate: (time.Now().Add(dur1).Format("2006-01-02T15:04:05.999999999Z07:00")),
			Complete:     false,
		},
		{
			To:           "dummyemail@dummy.dummy",
			Template:     "hello {{name}}",
			Data:         `{name:"test"}`,
			Subject:      "Test",
			ScheduleDate: (time.Now().Add(dur2).Format("2006-01-02T15:04:05.999999999Z07:00")),
			Complete:     false,
		},
		{
			To:           "dummyemail@dummy.dummy",
			Template:     "hello {{name}}",
			Data:         `{name:"test"}`,
			Subject:      "Test",
			ScheduleDate: (time.Now().Add(dur3).Format("2006-01-02T15:04:05.999999999Z07:00")),
			Complete:     true,
		},
	}
	for _, post := range postTable {
		fmt.Println(dbInstance)
		dbInstance.AddEmailRequest(&post)
	}

	if len(dbInstance.GetAllEmailSchedules()) != 3 {
		t.Errorf("Expected 2 scheduled posts found  %v", len(dbInstance.GetAllEmailSchedules()))
	}
}

func TestRecoverSchedules(t *testing.T) {
	dbInstance := GetInstance()
	recoverlist := dbInstance.RecoverSchedules()
	if len(recoverlist) != 2 {
		t.Errorf("Expected 2 recovery posts found %v", len(recoverlist))
	}
}

func TestRecoverEmailSchedules(t *testing.T) {
	dbInstance := GetInstance()
	recoverlist := dbInstance.RecoverEmailSchedules()
	if len(recoverlist) != 2 {
		t.Errorf("Expected 2 recovery posts found %v", len(recoverlist))
	}
}
