package actions

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	logger "sheddit/logger"
	scheduledb "sheddit/scheduleDatabase"
	session "sheddit/session"

	"github.com/mozillazg/request"
)

// BulkPost : Submits same link to different subreddits
func BulkPost(subreddits []string, link, title string, scheduleID int) {
	authToken := session.GetAuthToken()
	if authToken != nil {
		for _, subreddit := range subreddits {
			submitLink(link, title, subreddit, authToken.AccessToken)
		}
		scheduledb.UpdateStatus(scheduleID)
	}

}

func submitLink(link, title, subreddit, access_token string) {
	res, err := request.Post((os.Getenv("secure_api") + "/submit"), &request.Args{
		Client: new(http.Client),
		Headers: map[string]string{
			"Authorization": ("bearer " + access_token),
			"UserAgent":     os.Getenv("useragent"),
		},
		Data: map[string]string{
			"kind":  "link",
			"sr":    subreddit,
			"url":   link,
			"title": title,
		},
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
				logger.Log("Successfully posted link " + link + " in subreddit " + subreddit)
			} else {
				logger.Log("Error message " + string(bodyByte))
			}
		}
	}
}
