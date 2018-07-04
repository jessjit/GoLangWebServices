package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response, err := http.Get("https://developers.facebook.com/")
	if err != nil {
		fmt.Printf("the request could not be handled %s", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
	jsondata := map[string]string{"firstname": "Nick", "lastname": "Roboy"}
	jsonvalue, _ := json.Marshal(jsondata)
	request, _ := http.NewRequest("POST", "https://httpbin.org/post", bytes.NewBuffer(jsonvalue))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err = client.Do(request)
	//response, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonvalue))
	if err != nil {
		fmt.Printf("the request could not be handled %s", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}
