package actions

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	auth "sheddit/authentication"
	"sheddit/logger"
	"sheddit/types"
	"strings"

	"github.com/mozillazg/request"
)

func redditPost(post types.ScheduleRequest) bool {
	authToken := auth.GetAuthToken()
	subreddits := strings.Split(post.Subreddits, ",")
	if authToken != nil {
		for _, subreddit := range subreddits {
			submitPost(createData(post, subreddit), authToken.AccessToken, post.ScheduleID)
		}
		return true
	}
	return false
}

func submitPost(data map[string]string, access_token, scheduleID string) {
	res, err := request.Post((os.Getenv("secure_api") + "/submit"), &request.Args{
		Client: new(http.Client),
		Headers: map[string]string{
			"Authorization": ("bearer " + access_token),
			"UserAgent":     os.Getenv("useragent"),
		},
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}
	bodyByte, err := res.Content()
	var bodymap map[string]interface{}
	if err != nil {
		log.Print(err)
	} else {
		if err := json.Unmarshal(bodyByte, &bodymap); err != nil {
			fmt.Println(err)
		} else {
			success := bodymap["success"].(bool)
			if success {
				logger.Log("Successfully posted post with Schedule ID: " + scheduleID)
			} else {
				logger.Log("Error message " + string(bodyByte))
			}
		}
	}
}

func createData(post types.ScheduleRequest, subreddit string) map[string]string {

	if postType(post) == "link" {
		return map[string]string{
			"kind":  "link",
			"sr":    subreddit,
			"url":   post.Link,
			"title": post.Title,
		}
	}
	return map[string]string{
		"kind":  "self",
		"sr":    subreddit,
		"title": post.Title,
		"text":  post.Text,
	}
}

func postType(post types.ScheduleRequest) string {
	if post.Text != "" {
		return "text"
	}
	return "link"
}
