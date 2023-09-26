package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

var applicationStatus bool
var urls []string
var urlsProcessed int
var fullText string
var totalURLCount int
var wg sync.WaitGroup

func readURLs(statusChannel chan int, textChannel chan string) {
	defer wg.Done() // Decrement the WaitGroup when this function exits
	time.Sleep(time.Millisecond * 1)
	fmt.Println("Grabbing", len(urls), "urls")
	for i := 0; i < totalURLCount; i++ {
		resp, err := http.Get(urls[i])
		if err != nil {
			fmt.Println("Error getting URL")
			statusChannel <- 1
			continue // Continue to the next URL in case of an error
		}
		defer resp.Body.Close() // Close the response body when done with it
		text, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading body")
			statusChannel <- 1
			continue // Continue to the next URL in case of an error
		}
		textChannel <- string(text)
		statusChannel <- 0
		fmt.Println("Url", i, urls[i])
	}
}

func addToScrapedText(textChannel chan string) {
	defer wg.Done() // Decrement the WaitGroup when this function exits
	for {
		if applicationStatus == false {
			fmt.Println("Done")
			return // Exit the function when all URLs are processed
		}
		select {
		case tC := <-textChannel:
			fullText += tC
		}
	}
}

func evaluateStatus(statusChannel chan int) {
	defer wg.Done() // Decrement the WaitGroup when this function exits
	for {
		select {
		case status := <-statusChannel:
			fmt.Println("Case", urlsProcessed, totalURLCount, status)
			urlsProcessed++
			if status == 0 {
				fmt.Println("Got url")
			}
			if urlsProcessed == totalURLCount {
				fmt.Println("Read all top-level URLs")
				applicationStatus = false
				return // Exit the function when all URLs are processed
			}
		}
	}
}

func main() {
	applicationStatus = true
	statusChannel := make(chan int)
	textChannel := make(chan string)
	totalURLCount = 0

	urls = append(urls, "https://oguzhantasimaz.com/")
	urls = append(urls, "https://oguzhantasimaz.com/photography")
	urls = append(urls, "https://oguzhantasimaz.com/contact")

	fmt.Println("Starting spider")
	urlsProcessed = 0
	totalURLCount = len(urls)
	fmt.Println("Total URL count", totalURLCount)
	wg.Add(3) // Add 3 goroutines to the WaitGroup
	go readURLs(statusChannel, textChannel)
	go evaluateStatus(statusChannel)
	go addToScrapedText(textChannel)

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println(fullText)
	fmt.Println("Done!")
}
