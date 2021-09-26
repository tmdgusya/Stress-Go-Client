package main

import (
	"fmt"

	customer "stress-go/src/thread_customer"
)

const PERIOD = 600;

func main() {
	jobs := make(chan int, 500)
	reports := make(chan bool);
	
	testReport := customer.ReportFactory()

	for w := 1; w <= PERIOD; w++ {
		go customer.CustomerFactory(2).ConnectUser(w, jobs, reports)
	}

	for j := 1; j <= PERIOD; j++ {
		jobs <- j
	}

	close(jobs)

	for r := 1; r <= PERIOD; r++ {
		result := <- reports

		if result {
			testReport.SuccessTestResult()
		} else {
			testReport.FailTestResult()
		}
	}

	fmt.Println("Success Count : ", testReport.GetSuccessCount(), " Fail Count : ", testReport.GetFailCount())
}
