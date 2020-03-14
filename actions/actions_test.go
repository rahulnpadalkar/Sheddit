package actions

import (
	"sheddit/types"
	"testing"
)

//TestBulkPost : test for BulkPost in actions package
func TestBulkPost(t *testing.T) {

	postTable := []types.ScheduleRequest{
		{
			Subreddits: "test,sandboxtest",
			Link:       "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
			Title:      "Must watch!",
			Text:       "",
			ScheduleID: 1,
		},
		{
			Subreddits: "test,sandboxtest",
			Link:       "https://www.youtube.com/watch?v=Fkk9DI-8el4",
			Title:      "testPost2",
			Text:       "",
			ScheduleID: 2,
		},
		{
			Subreddits: "test,sandboxtest",
			Text:       "Hello!",
			Title:      "testPost2",
			Link:       "",
			ScheduleID: 3,
		},
	}
	for _, v := range postTable {
		success := BulkPost(v)
		if !success {
			t.Errorf("Failed for : subreddits: %v, link: %v and title %v", v.Subreddits, v.Link, v.Title)
		}
	}
}
