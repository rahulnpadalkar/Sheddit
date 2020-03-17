package types

// ScheduleRequest : Represents structure of a scheduled post
type ScheduleRequest struct {
	Subreddits   string `binding:"required"`
	Title        string `binding:"required"`
	Text         string `required_without:"Link"`
	Link         string `required_without:"Text"`
	ScheduleDate string `binding:"required"`
	ScheduleID   int
	Complete     bool
	Provider     string `binding:"required"`
}

// TestSchedulePost : Represents a test structur used for testing BulkPost
type TestSchedulePost struct {
	Subreddits []string
	Link       string
	Title      string
}
