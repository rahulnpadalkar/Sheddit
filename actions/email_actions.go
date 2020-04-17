package actions

import (
	"encoding/json"
	"os"
	"sheddit/logger"
	"sheddit/types"
	"strconv"
	"strings"

	handlebar "github.com/aymerick/raymond"
	_ "github.com/joho/godotenv/autoload"
	"gopkg.in/gomail.v2"
)

// SendEmail : Send email using plain SMTP auth
func SendEmail(emailRequest *types.EmailRequest) bool {
	recipients := strings.Split(emailRequest.To, ",")
	dataArray, err := parseData(emailRequest.Data)
	if err != nil {
		return false
	}
	port, err := strconv.Atoi(os.Getenv("smtp_port"))
	if err != nil {
		logger.Log(err.Error())
		return false
	}
	dialier := gomail.NewDialer(os.Getenv("smtp_server"), port, os.Getenv("smtp_username"), os.Getenv("smtp_password"))

	for i := 0; i < len(recipients); i++ {
		mail := gomail.NewMessage()
		compiled, err := compileTemplate(emailRequest.Template, dataArray[i].(map[string]interface{}))
		if err != nil {
			return false
		}
		mail.SetHeader("To", recipients[i])
		mail.SetHeader("From", os.Getenv("smtp_sender"))
		mail.SetHeader("Subject", emailRequest.Subject)
		mail.SetBody("text/html", compiled)
		if err := dialier.DialAndSend(mail); err != nil {
			logger.Log(err.Error())
			return false
		}
	}
	return true
}

func parseData(data string) ([]interface{}, error) {
	var unstructData []interface{}
	err := json.Unmarshal([]byte(data), &unstructData)
	if err != nil {
		logger.Log(err.Error())
		return nil, err
	}
	return unstructData, nil
}

func compileTemplate(template string, data map[string]interface{}) (string, error) {
	output, err := handlebar.Render(template, data)
	if err != nil {
		logger.Log(err.Error())
		return "", err
	}
	return output, nil
}
