package actions

import (
	"fmt"
	auth "sheddit/authentication"
	"sheddit/logger"
	"sheddit/types"
)

func twitterTweet(post types.ScheduleRequest) (bool, int64) {
	client := auth.GetClient()
	tweet, _, err := client.Statuses.Update(post.Text, nil)
	if err != nil {
		logger.Log("Error while tweeting: " + post.Text + err.Error())
		return false, 0
	}
	return true, tweet.ID
}

func deleteTestTweets(ids []int64) {
	client := auth.GetClient()
	for _, v := range ids {
		_, _, err := client.Statuses.Destroy(v, nil)
		if err != nil {
			fmt.Printf("Error deleting test tweets: %v", err)
		}
	}
}
