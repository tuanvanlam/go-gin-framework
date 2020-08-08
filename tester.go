package main

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "localhost:8080/videos"
	method := "GET"

	payload := strings.NewReader("{\n    \"title\": \"title 1\",\n    \"description\": \"desc 1\",\n    \"url\": \"abc.com\"\n}")

	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Authorization", "Basic YWRtaW46YWRtaW4=")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}