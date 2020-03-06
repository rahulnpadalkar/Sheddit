package session

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/mozillazg/request"
)

type AuthToken struct {
	AccessToken string
	Expiry      time.Time
}

type authResponse struct {
	Access_token string
	Expiry       int64
	Scope        string
}

func GetAuthToken() *AuthToken {
	c := new(http.Client)
	auth := authResponse{}
	req := request.NewRequest(c)
	req.BasicAuth = request.BasicAuth{os.Getenv("clientid"), os.Getenv("clientsecret")}
	req.Data = map[string]string{
		"grant_type": "password",
		"username":   os.Getenv("username"),
		"password":   os.Getenv("password"),
	}
	req.Headers = map[string]string{
		"User-Agent": os.Getenv("useragent"),
	}
	res, err := req.Post("https://www.reddit.com/api/v1/access_token")

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		err := json.NewDecoder(res.Body).Decode(&auth)
		if err != nil {
			log.Fatal(err)
		}
		duration, err := time.ParseDuration(string(strconv.FormatInt(auth.Expiry, 10)) + "s")
		if err != nil {
			log.Fatal(err)
		}
		expiryTime := time.Now().Add(duration)

		return &AuthToken{
			Expiry:      expiryTime,
			AccessToken: auth.Access_token,
		}
	}
	return nil
}
