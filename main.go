package main

import (
	"fmt"
	"log"
	"net/url"

	stomp "github.com/drawdy/stomp-ws-go"
	"github.com/gorilla/websocket"
)

const PERIOD = 10;

// func main() {
// 	jobs := make(chan int, 500)
// 	reports := make(chan bool);
	
// 	testReport := customer.ReportFactory()

// 	for w := 1; w <= PERIOD; w++ {
// 		go customer.CustomerFactory(2).ConnectUser(w, jobs, reports)
// 	}

// 	for j := 1; j <= PERIOD; j++ {
// 		jobs <- j
// 	}

// 	close(jobs)

// 	for r := 1; r <= PERIOD; r++ {
// 		result := <- reports

// 		if result {
// 			testReport.SuccessTestResult()
// 		} else {
// 			testReport.FailTestResult()
// 		}
// 	}

// 	fmt.Println("Success Count : ", testReport.GetSuccessCount(), " Fail Count : ", testReport.GetFailCount())
// }

func main() {

	u := url.URL{
		Scheme: "ws",
		Host:   "127.0.0.1:8761",
		Path:   "/stomp",
	}

	log.Println(u.String())

	conn, resp, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatalf("couldn't connect to %v: %v", u.String(), err)
	}
	log.Printf("response status: %v\n", resp.Status)
	log.Print("Websocket connection succeeded.")

	h := stomp.Headers{
		stomp.HK_ACCEPT_VERSION, "1.2,1.1,1.0",
		stomp.HK_HEART_BEAT, "3000,3000",
		stomp.HK_HOST, u.Host,
		"token", "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJpc3MiOiJjbGVhbmluZ19sYWIiLCJleHAiOjE2MzI3MzI3MTYsInRva2VuX2RhdGEiOnsiYXBwSWQiOjAsImFsbG93T3JpZ2lucyI6bnVsbH19.w1ct3ZC7impouFSOAeLR2mJBMF67zmpb1VKys2ce730CFA7t1uO-nmPThQ1XiEFvQ1GpX6FDqloFMDi96JDJM7WinyDYPxvOVjXkySRu1S3_d4qVe_iyNI6G4w31mJ4RJFxSRi8sj3PugeSGadC0I75L_t7o2vQOnPXDOp1KrymYHWaratdCpNYxKL5xC12jqW6gUGlXc8FvyrflDriClfawjzD6ni38-SLixaDLrtWML5claI28BKJQqu3Gvb_duOWhBnLgwXQQfveMxTbwllqBmpc10SdBOS6Tp-bUjRE0usRpnhnRG_Nhg8KptKQUPg4S6ggNg4jVBp7YZ3xtjg",
		"host",  "development-app-0"}
	sc, err := stomp.ConnectOverWS(conn, h)
	if err != nil {
		log.Fatalf("couldn't create stomp connection: %v", err)
	}

	mdCh, err := sc.Subscribe(stomp.Headers{
		stomp.HK_DESTINATION, "/topic/chat",
		stomp.HK_ID, stomp.Uuid(),
	})
	if err != nil {
		log.Fatalf("failed to suscribe greeting message: %v", err)
	}

	err = sc.Send(stomp.Headers{
		stomp.HK_DESTINATION, "/topic/chat",
		stomp.HK_ID, stomp.Uuid(),
	}, "hello STOMP!")
	if err != nil {
		log.Fatalf("failed to send greeting message: %v", err)
	}

	md := <-mdCh
	if md.Error != nil {
		log.Fatalf("receive greeting message caught error: %v", md.Error)
	}

	fmt.Printf("----> receive new message: %v\n", md.Message.BodyString())

	err = sc.Disconnect(stomp.NoDiscReceipt)
	if err != nil {
		log.Fatalf("failed to disconnect: %v", err)
	}

	log.Print("Disconnected.")
}
