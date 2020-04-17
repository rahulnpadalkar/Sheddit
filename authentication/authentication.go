package authentication

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	_ "github.com/joho/godotenv/autoload"
	"github.com/mozillazg/request"
)

type AuthToken struct {
	AccessToken string
	Expiry      time.Time
}

type authResponse struct {
	Access_token string
	Expires_in   int64
	Scope        string
}

var client *twitter.Client
var authToken AuthToken

// GetAuthToken : Get auth token for making calls to the api.
func GetAuthToken() *AuthToken {
	if !authTokenExpired() {
		return &authToken
	}
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
		duration, err := time.ParseDuration(string(strconv.FormatInt(auth.Expires_in, 10)) + "s")
		if err != nil {
			log.Fatal(err)
		}
		expiryTime := time.Now().Add(duration)
		authToken = AuthToken{
			Expiry:      expiryTime,
			AccessToken: auth.Access_token,
		}
		return &authToken
	}
	return nil
}

func authTokenExpired() bool {
	if authToken.Expiry.After(time.Now()) {
		return false
	}
	return true
}

// GetClient : Get Twitter Client
func GetClient() *twitter.Client {
	if client != nil {
		return client
	}
	config := oauth1.NewConfig(os.Getenv("t_consumerkey"), os.Getenv("t_consumersecret"))
	token := oauth1.NewToken(os.Getenv("t_accesstoken"), os.Getenv("t_accessecret"))
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)
	return client
}
