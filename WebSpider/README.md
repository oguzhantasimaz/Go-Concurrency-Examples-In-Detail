# Concurrent Web Scraping and Data Processing in Go

This Go program showcases concurrent web scraping and data processing. It demonstrates how to scrape multiple web pages simultaneously, process the retrieved data concurrently, and collect the results into a single dataset.

## Overview

In this program, we have three main components:

1. **URL List**: A list of URLs to scrape is defined in the `urls` slice.

2. **Goroutines**: We use goroutines to perform concurrent web scraping and data processing. Specifically, there are three types of goroutines:
    - `readURLs`: This goroutine sends HTTP requests to the specified URLs and reads their content. It communicates the results to other goroutines using channels.
    - `evaluateStatus`: This goroutine evaluates the status of the scraping process and keeps track of the number of URLs processed.
    - `addToScrapedText`: This goroutine collects and processes the retrieved text data.

3. **Channels**: Channels are used for communication and synchronization between goroutines. Two types of channels are employed:
    - `statusChannel`: Used to signal the success or failure of URL retrieval.
    - `textChannel`: Used to pass the retrieved text data.

## Code Explanation

- `readURLs`: This function sends HTTP requests to the specified URLs, reads their content, and communicates the results to other goroutines using channels. It handles errors and continues to the next URL in case of failures.

- `addToScrapedText`: This function collects and processes the retrieved text data and stores it in the `fullText` variable. It continuously processes data as long as the application status is active.

- `evaluateStatus`: This function evaluates the status of URL retrieval, counts the number of processed URLs, and terminates the program when all URLs have been processed.
