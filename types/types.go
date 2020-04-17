package types

// ScheduleRequest : Represents structure of a scheduled post
type ScheduleRequest struct {
	Subreddits   string
	Title        string
	Text         string
	Link         string
	ScheduleDate string `binding:"required"`
	ScheduleID   string `pg:"pk"`
	Complete     bool   `sql:",notnull"`
	Provider     string `binding:"required,oneof=reddit Reddit twitter Twitter"`
}

// EmailRequest : Represents structure of a email request
type EmailRequest struct {
	To           string `binding:"required"`
	Template     string `binding:"required"`
	Data         string `binding:"required"`
	ScheduleDate string `binding:"required"`
	Subject      string `binding:"required"`
	ScheduleID   string
	Complete     bool `sql:",notnull"`
}
