package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request Failed")

type result struct {
	url    string
	status string
}

func main() {
	results := map[string]string{}
	channel := make(chan result)
	urls := []string{
		"https://google.com",
		"https://naver.com",
		"https://facebook.com",
		"https://reddit.com",
	}

	for _, url := range urls {
		go hitURL(url, channel)

	}
	for i := 0; i < len(urls); i++ {
		result := <-channel
		results[result.url] = result.status
	}
	for url, status := range results {
		fmt.Println("URL is ", url, " Status is ", status)
	}

}

func hitURL(url string, channel chan result) { //channel chan <- result => 보내기만 가능, 받는건 불가능
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		channel <- result{url: url, status: "Failed"}
	}
	channel <- result{url: url, status: "Okay"}
}

// Goroutines -> 함수를 동시에 실행시키는 함수
