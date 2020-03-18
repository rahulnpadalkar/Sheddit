package types

// ScheduleRequest : Represents structure of a scheduled post
type ScheduleRequest struct {
	Subreddits   string
	Title        string
	Text         string
	Link         string
	ScheduleDate string `binding:"required"`
	ScheduleID   int
	Complete     bool
	Provider     string `binding:"required,oneof=reddit Reddit twitter Twitter"`
}

// TestSchedulePost : Represents a test structur used for testing BulkPost
type TestSchedulePost struct {
	Subreddits []string
	Link       string
	Title      string
}
