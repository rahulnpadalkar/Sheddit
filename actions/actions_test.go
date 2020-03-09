package actions

import (
	"sheddit/types"
	"testing"
)

//TestBulkPost : test for BulkPost in actions package
func TestBulkPost(t *testing.T) {

	postTable := []types.TestSchedulePost{
		{
			Subreddits: []string{"test", "sandboxtest"},
			Link:       "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
			Title:      "Must watch!",
		},
		{
			Subreddits: []string{"test", "sandboxtest"},
			Link:       "https://www.youtube.com/watch?v=Fkk9DI-8el4",
			Title:      "testPost2",
		},
	}
	for _, v := range postTable {
		success := BulkPost(v.Subreddits, v.Link, v.Title)
		if !success {
			t.Errorf("Failed for : subreddits: %v, link: %v and title %v", v.Subreddits, v.Link, v.Title)
		}
	}
}
