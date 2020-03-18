package actions

import (
	"sheddit/types"
	"strings"
)

// BulkPost : Submits same link to different subreddits
func BulkPost(post types.ScheduleRequest) (bool, int64) {
	if strings.EqualFold(post.Provider, "reddit") {
		return redditPost(post), 0
	} else if strings.EqualFold(post.Provider, "twitter") {
		success, id := twitterTweet(post)
		return success, id
	}
	return false, 0
}
