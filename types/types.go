package types

type ScheduleRequest struct {
	Subreddits   string `binding:"required"`
	Title        string `binding:"required"`
	Link         string
	ScheduleDate string `binding:"required"`
	ScheduleID   int
	Complete     bool
}

type TestSchedulePost struct {
	Subreddits []string
	Link       string
	Title      string
}
