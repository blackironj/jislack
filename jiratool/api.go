package jiratool

import (
	"encoding/base64"
	"fmt"
	"log"

	resty "github.com/go-resty/resty/v2"
)

var jiraClient *resty.Client

func Init(username, apiToken, baseUrl string) {
	auth := base64.StdEncoding.EncodeToString([]byte(username + ":" + apiToken))
	fmt.Println(auth)

	client := resty.New()
	client.SetHeader("Authorization", "Basic "+auth)
	client.SetHeader("Content-Type", "application/json")
	client.SetHeader("Accept", "application/json")

	client.SetHostURL(baseUrl)

	jiraClient = client
}

func GetUserAccountIdByEmail(email string) string {
	resp, err := jiraClient.R().SetQueryParams(map[string]string{
		"query": email,
	}).Get("/rest/api/3/user/search")
	if err != nil {
		log.Println(err)
		return ""
	}
	fmt.Println("user body : ", string(resp.Body()))
	return ""
}

func CreateIssue(projKey, issueType, summary, description string) {
	body := map[string]interface{}{
		"fields": map[string]interface{}{
			"project": map[string]interface{}{
				"key": projKey,
			},
			"summary":     summary,
			"description": description,
			"issuetype": map[string]interface{}{
				"name": issueType,
			},
		},
	}
	resp, err := jiraClient.R().SetBody(body).Post("/rest/api/3/issue")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(resp.Body()))

}
