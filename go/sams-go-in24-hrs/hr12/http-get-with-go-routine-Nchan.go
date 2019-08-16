package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var flag int = 0

func getSites(url string, stopChan chan string) {
	start := time.Now()
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	elapsed := time.Since(start).Seconds()
	fmt.Printf("Elaspsed time : %v sec for url %s\n", elapsed, url)
	flag++
	if flag == 3 {
		stopChan <- "stop"
	}
}

func main() {
	urls := make([]string, 3)
	urls[0] = "https://www.usa.gov/"
	urls[1] = "https://www.gov.uk/"
	urls[2] = "https://www.india.gov.in/"

	stopChan := make(chan string)

	for _, url := range urls {
		go getSites(url, stopChan)
	}

	select {
	case <-stopChan:
		fmt.Println("Got stop msg.")
		return
	}
}
