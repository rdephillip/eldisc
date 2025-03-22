package main

import (
	"io"
	"log"
	"net/http"

	"github.com/tidwall/gjson"
)

func participantInfo() (goal, raised, avatar, donate string) {
	// Call the configured API and retrieve the information to reply with a goal status
	ep := "participants/" + Pid
	req, err := http.NewRequest("GET", Url+ep, nil)
	if err != nil {
		log.Println(err.Error())
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err.Error())
	}

	defer res.Body.Close()
	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		log.Println(err.Error())
	}

	goal = gjson.Get(string(body), "fundraisingGoal").String()
	raised = gjson.Get(string(body), "sumDonations").String()
	avatar = gjson.Get(string(body), "avatarImageURL").String()
	donate = gjson.Get(string(body), "links.page").String()

	return
}

func participantDonations() (body []byte) {
	// API call for configured participants donations
	ep := "participants/" + Pid + "/donations"

	req, err := http.NewRequest("GET", Url+ep, nil)
	if err != nil {
		log.Println(err.Error())
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err.Error())
	}

	defer res.Body.Close()
	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		log.Println(err.Error())
	}

	return
}
