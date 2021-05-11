package main

import (
	"fmt"
	"log"
	"net/http"
)

var loginURL string = "https://talk.tmax.co.kr/login.do"

func main() {
	fmt.Println("Starting login...")
}

func getPages() int {
	res, err := http.Get(loginURL)
	checkError(err)
	checkStatusCode(res)
	return 0
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkStatusCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with status: ", res.StatusCode)
	}
}
