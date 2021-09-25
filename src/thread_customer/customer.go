package threadcustomer

import (
	"fmt"
	"stress-go/src/connection"
	"time"
)

type Customer struct {
	period int
}

func (c *Customer) connect(customer_id int, jobs <- chan string, result chan <- bool) {
	for job := range jobs {
		fmt.Println("Customer", customer_id ,"Connecting ....", job);
		time.Sleep(time.Duration(int(time.Second) * c.period));

		_, err := connection.ConnectFactory();

		if err != nil {
			result <- false
		}
		if err == nil {
			result <- true
		} 
	}
}