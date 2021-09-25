package main

import (
	customer "stress-go/src/thread_customer"
)

func main() {
	jobs := make(chan int, 500)
	reports := make(chan bool, 500);

	for w := 1; w <= 500; w++ {
		go customer.CustomerFactory(2).ConnectUser(w, jobs, reports)
	}

	for j := 1; j <= 500; j++ {
		jobs <- j
	}

	close(jobs)

	for r := 1; r <= 500; r++ {
		<- reports
	}
}
