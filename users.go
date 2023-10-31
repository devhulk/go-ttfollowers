package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
  "os"
)

var baseURL = os.Getenv("TT_SECRET_URL")

type User struct {
	SecUID   string    `json:"sec_uid"`
	UniqueID string    `json:"uniqueId"`
	Name     string    `json:"name"`
	Data     *UserData `json:"userData"`
}

type UserData struct {
	AvatarThumb    string `json:"avatarThumb"`
	FollowerCount  uint   `json:"followerCount"`
	FollowingCount uint   `json:"followingCount"`
	HeartCount     uint   `json:"heartCount"`
	VideoCount     uint   `json:"videoCount"`
	Message        string `json:"message"`
	Status         string `json:"status"`
}

type Users struct {
	Users []User `json:"users"`
}

func (u *User) getUserInfo() *UserData {

	var data UserData
	tries := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, v := range tries {
		time.Sleep(1 * time.Second)
		request, err := http.NewRequest("GET", baseURL, nil)
		if err != nil {
			log.Printf("REQUEST INIT ERROR: %v", err)
		}

		request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36")

		q := request.URL.Query()
		q.Add("sec_user_id", u.SecUID)

		request.URL.RawQuery = q.Encode()

		response, err := http.Get(request.URL.String())
		if err != nil {
			fmt.Printf("ERROR WITH RESPONSE: %v", err)
			break
		}

		defer response.Body.Close()

		err = json.NewDecoder(response.Body).Decode(&data)
		if err != nil {
			fmt.Printf("ERROR Parsing Response Body: %v", err)
		}

		if data.Status == "error" {
			log.Printf("%s data fetch FAILED on try %v", u.Name, v)
			continue
		}

		if data.Status == "success" {
			log.Printf("%s data fetch SUCCESS on try %v", u.Name, v)
			break
		}

	}

	u.Data = &data
	return &data

}

// Util / Debug Function
func readBodyString(response *http.Response) string {

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("ERROR reading body: %v", body)
	}

	res := string(body)
	return res

}
