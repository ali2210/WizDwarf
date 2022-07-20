package isconnect

import (
	"net/http"
	"time"
)

//IsOnline is a function to check wheter an internet connection
//return bool
func IsOnline() bool {
	timeout := time.Duration(5000 * time.Millisecond)
	client := http.Client{
		Timeout: timeout,
	}
	//default url to check connection is http://google.com
	_, err := client.Get("https://google.com")

	if err != nil {
		return false
	}

	return true
}

//IsReachable is a function to check whether a reacheable network
//return bool and status/error message
func IsReachable(url string) (bool, string) {
	timeout := time.Duration(5000 * time.Millisecond)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(url)

	if err != nil {
		return false, err.Error()
	}

	return true, resp.Status
}
