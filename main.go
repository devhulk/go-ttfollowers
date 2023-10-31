package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

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

}
