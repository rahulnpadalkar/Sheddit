package scheduledatabase

import (
	"os"
	"sheddit/types"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	InitializeDB()
	deleteBucket(os.Getenv("bucketname"))
	os.Exit(m.Run())
}
func TestAddToSchedule(t *testing.T) {
	dur1, _ := time.ParseDuration("180s")
	dur2, _ := time.ParseDuration("180s")
	dur3, _ := time.ParseDuration("-30s")
	postTable := []types.ScheduleRequest{
		{
			Subreddits:   "test,sandboxtest",
			Title:        "Test Post",
			Link:         "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
			ScheduleDate: (time.Now().Add(dur1).Format("2006-01-02T15:04:05.999999999Z07:00")),
		},
		{
			Subreddits:   "test,sandboxtest",
			Title:        "Test Post",
			Link:         "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
			ScheduleDate: (time.Now().Add(dur2).Format("2006-01-02T15:04:05.999999999Z07:00")),
		},
		{
			Subreddits:   "test,sandboxtest",
			Title:        "Test Post",
			Link:         "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
			ScheduleDate: (time.Now().Add(dur3).Format("2006-01-02T15:04:05.999999999Z07:00")),
		},
	}
	for _, post := range postTable {
		AddToSchedule(&post)
	}

	if len(GetAllSchedules()) != 3 {
		t.Errorf("Expected 2 scheduled posts found  %v", len(GetAllSchedules()))
	}
}

func TestRecoverSchedules(t *testing.T) {
	recoverlist := RecoverSchedules()

	if len(recoverlist) != 2 {
		t.Errorf("Expected 2 recovery posts found %v", len(recoverlist))
	}
}
