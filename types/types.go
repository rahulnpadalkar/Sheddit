package types

// ScheduleRequest : Represents structure of a scheduled post
type ScheduleRequest struct {
	Subreddits   string `binding:"required"`
	Title        string `binding:"required"`
	Link         string `binding:"required"`
	ScheduleDate string `binding:"required"`
	ScheduleID   int
	Complete     bool
}

// TestSchedulePost : Represents a test structur used for testing BulkPost
type TestSchedulePost struct {
	Subreddits []string
	Link       string
	Title      string
}
