package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type RandomUser struct {
Results []UserResult
}

type UserResult struct {
	Name UserName `json:"username"`
	Email string `json:"email"`
	Picture UserPicture `json:"picture"`
}
type UserName struct {
	Title string
	First string
	Last string
}
type UserPicture struct {
	Largest string
	Medium string
	Thumbnail string
}


var client  *http.Client
type CatFact struct {
	Fact string `json:"fact"`
	Length int `json:"length"`
}
func GetCatFact() {
	url:= "https://catfact.ninja/fact"
	var catFact CatFact
	err := GetJson(url, &catFact)
	if err != nil {
		fmt.Printf("error getting cat fact: %s\n", err.Error())
		return
	} else {
		fmt.Printf("A interesting cat fact: %s\n", catFact.Fact)
	}
}

func GetRandomUser() {
	url:= "https://randomuser.me/api/?inc=name,email,picture"
	var randomuser RandomUser
	err := GetJson(url, &randomuser)
	if err != nil {
		fmt.Printf("error getting json: %s\n", err.Error())
		return
	} else {
		fmt.Printf("User: %s %s %s\nEmail: %s\nThumbnail: %s\n",
			randomuser.Results[0].Name.Title,
			randomuser.Results[0].Name.First,
			randomuser.Results[0].Name.Last,
			randomuser.Results[0].Email,
			randomuser.Results[0].Picture.Thumbnail,
		)
	}
}

func main() {
	client = &http.Client{Timeout: 10 * time.Second}
	GetCatFact()
	GetRandomUser()
}


func GetJson(url string, target interface{}) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return json.NewDecoder(response.Body).Decode(target)
}