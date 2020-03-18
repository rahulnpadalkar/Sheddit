package actions

import (
	"sheddit/types"
	"testing"
)

//TestBulkPost : test for BulkPost in actions package
func TestBulkPost(t *testing.T) {
	var tweetIds []int64
	postTable := []types.ScheduleRequest{
		{
			Subreddits: "test,sandboxtest",
			Link:       "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
			Title:      "Must watch!",
			Text:       "",
			ScheduleID: 1,
			Provider:   "Reddit",
		},
		{
			Text:       "Test Tweet!",
			ScheduleID: 4,
			Provider:   "Twitter",
		},
		{
			Subreddits: "test,sandboxtest",
			Text:       "Hello!",
			Title:      "testPost2",
			Link:       "",
			ScheduleID: 5,
			Provider:   "Twitter",
		},
	}
	for _, v := range postTable {
		success, id := BulkPost(v)
		if id != 0 {
			tweetIds = append(tweetIds, id)
		}
		if !success {
			t.Errorf("Failed for :%v", v)
		}
	}
	deleteTestTweets(tweetIds)
}
