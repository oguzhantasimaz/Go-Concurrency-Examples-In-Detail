package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var balance_w_race int
var transactionNo_w_race int

func main() {
	rand.Seed(time.Now().Unix())
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	tranChan := make(chan bool)
	balance_w_race = 1000
	transactionNo_w_race = 0
	fmt.Println("Starting balance_w_race: $", balance_w_race)
	wg.Add(1)
	for i := 0; i < 100; i++ {
		go func(ii int, trChan chan bool) {
			transactionAmount := rand.Intn(25)
			Transaction(transactionAmount)
			if ii == 99 {
				trChan <- true
			}
		}(i, tranChan)
	}
	go Transaction(0)
	select {
	case <-tranChan:
		fmt.Println("Transactions finished")
		wg.Done()
	}
	wg.Wait()
	close(tranChan)
	fmt.Println("Final balance_w_race: $", balance_w_race)
}

func Transaction(amt int) bool {
	fmt.Println("\tBefore balance_w_race $", balance_w_race)
	approved := false
	if (balance_w_race - amt) < 0 {
		approved = false
	} else {
		approved = true
		balance_w_race = balance_w_race - amt
	}
	approvedText := "declined"
	if approved == true {
		approvedText = "approved"
	}
	transactionNo_w_race = transactionNo_w_race + 1
	fmt.Println(transactionNo_w_race, "Transaction for $", amt, approvedText)
	fmt.Println("\tRemaining balance_w_race $", balance_w_race)
	return approved
}
