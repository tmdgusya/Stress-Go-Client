package threadcustomer

import (
	"fmt"
	"log"
	"stress-go/src/connection"
	"time"

	stomp "github.com/drawdy/stomp-ws-go"
)

type Customer struct {
	period int
}

func (c *Customer) ConnectUser(customer_id int, jobs <- chan int, result chan <- bool) {
	for job := range jobs {
		fmt.Println("Customer", customer_id ,"Connecting ....", job);
		
		sc, err := connection.ConnectFactory();

		if err != nil {
			result <- false
			return
		}

		time.Sleep(time.Duration(int(time.Second) * c.period));

		ch, err := sc.Subscribe(stomp.Headers{
			stomp.HK_DESTINATION, "/topic/chat",
			stomp.HK_ID, stomp.Uuid(),
		})
		if err != nil {
			log.Print("failed to suscribe greeting message: %v", err)
		}
	
		// err = sc.Send(stomp.Headers{
		// 	stomp.HK_DESTINATION, "/topic/chat",
		// 	stomp.HK_ID, stomp.Uuid(),
		// }, "hello STOMP!")
		// if err != nil {
		// 	log.Print("failed to send greeting message: %v", err)
		// }
	
		md := <-ch
		if md.Error != nil {
			log.Print("receive greeting message caught error: %v", md.Error)
		}
	
		fmt.Printf("----> receive new message: %v\n", md.Message.BodyString())

		if err != nil {
			result <- false
		}
		if ch != nil {
			result <- true
		} 
	}
}

func CustomerFactory(period int) *Customer {
	return &Customer{period: period}
}