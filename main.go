package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

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

func main() {
	client = &http.Client{Timeout: 10 * time.Second}
	GetCatFact()
}


func GetJson(url string, target interface{}) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return json.NewDecoder(response.Body).Decode(target)
}