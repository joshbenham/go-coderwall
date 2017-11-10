package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/a8m/djson"
)

// CallAPI parses urls into valid JSON
func CallAPI(url string) (interface{}, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	response, err := client.Get(url)
	if err != nil {
		return "", StyleError(fmt.Sprintf("Could not Parse URL: %s", err))
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", StyleError(fmt.Sprintf("Could not Convert to Byte: %s", err))
	}

	return djson.Decode(body)
}
