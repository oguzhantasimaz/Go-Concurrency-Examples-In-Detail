package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID  int
	Msg string
}

func worker(id int, jobs <-chan Job, results chan<- Job, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing Job %d: %s\n", id, job.ID, job.Msg)
		//Simulate processing time
		time.Sleep(1 * time.Second)
		job.Msg = "Processed by Worker " + fmt.Sprintf("%d", id)
		results <- job
	}
}

func main() {
	numWorkers := 3
	numJobs := 10

	jobs := make(chan Job, numJobs)
	results := make(chan Job, numJobs)
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	for i := 1; i <= numJobs; i++ {
		job := Job{ID: i, Msg: fmt.Sprintf("Job %d", i)}
		jobs <- job
	}

	close(jobs)
	wg.Wait()
	close(results)
	for result := range results {
		fmt.Printf("Result from Job %d: %s\n", result.ID, result.Msg)
	}
}
