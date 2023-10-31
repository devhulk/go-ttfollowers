package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
)

var slackid string

func minus(a, b int) int {
	return a - b
}

func main() {
	slackid = os.Getenv("SLACK_SERVICE")

	usersFile, err := os.Open(os.Getenv("USER_FILE_NAME"))

	if err != nil {
		log.Printf("USERFILE LOAD ERR: %v", err)
	}

	defer usersFile.Close()

	byteValue, err := ioutil.ReadAll(usersFile)

	if err != nil {
		log.Printf("FILE CONVERSION FAILED: %v", err)
	}

	var users Users
	var userData []UserData

	json.Unmarshal(byteValue, &users)

	for i := 0; i < len(users.Users); i++ {
		currentUser := &users.Users[i]
		data := currentUser.getUserInfo()

		fmt.Printf("Username: %v \n", currentUser.Name)
		fmt.Printf("Follower Count: %v \n", currentUser.Data.FollowerCount)
		fmt.Printf("Following Count: %v \n", currentUser.Data.FollowingCount)
		fmt.Printf("Heart Count: %v \n", currentUser.Data.HeartCount)
		fmt.Printf("Video Count: %v \n", currentUser.Data.VideoCount)

		userData = append(userData, *data)
	}

	fmt.Println(userData)

	// Slack test

	var funcMap = template.FuncMap{
		"minus": minus,
	}
	tmpl, err := template.New("slack.tmpl.json").Funcs(funcMap).ParseFiles("slack.tmpl.json")
	if err != nil {
		fmt.Println(err)
	}
	f, err := os.Create("slack_request.json")
	if err != nil {
		log.Println("create file: ", err)
		return
	}
	err = tmpl.Execute(f, userData)
	if err != nil {
		fmt.Println(err)
	}

	slackFile, err := os.Open("slack_request.json")
	if err != nil {
		log.Printf("SLACK FILE LOAD ERR: %v", err)
	}

	defer slackFile.Close()

	byteSlackValue, err := ioutil.ReadAll(slackFile)
	if err != nil {
		fmt.Println(err)
	}

	var result map[string]interface{}
	err2 := json.Unmarshal([]byte(byteSlackValue), &result)
	if err2 != nil {
		fmt.Println(err2)
	}

	//fmt.Println(slackid)
	fmt.Println(string(byteSlackValue))
	req, err := http.NewRequest(http.MethodPost, slackid, bytes.NewBuffer(byteSlackValue))
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println((err))
	}

	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

}
