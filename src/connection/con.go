package connection

import (
	"log"
	"net/http"
	"net/url"

	stomp "github.com/drawdy/stomp-ws-go"
	stompws "github.com/drawdy/stomp-ws-go"
	"github.com/gorilla/websocket"
)


func ConnectFactory() (stompws.STOMPConnector, error) {
	u := url.URL{
		Scheme: "ws",
		Host:   "127.0.0.1:8761",
		Path:   "/stomp",
	}

	hh := http.Header{}

	hh.Set("origin", "**")

	log.Println(u.String())

	conn, resp, err := websocket.DefaultDialer.Dial(u.String(), hh)
	if err != nil {
		log.Print("couldn't connect to %v: %v", u.String(), err)
		return nil, err
	}
	log.Printf("response status: %v\n", resp.Status)
	log.Print("Websocket connection succeeded.")

	h := stomp.Headers{
		stomp.HK_ACCEPT_VERSION, "1.2,1.1,1.0",
		stomp.HK_HEART_BEAT, "3000,3000",
		stomp.HK_HOST, "development-app-0",
		"token", "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJpc3MiOiJjbGVhbmluZ19sYWIiLCJleHAiOjE2MzI3MzI3MTYsInRva2VuX2RhdGEiOnsiYXBwSWQiOjAsImFsbG93T3JpZ2lucyI6bnVsbH19.w1ct3ZC7impouFSOAeLR2mJBMF67zmpb1VKys2ce730CFA7t1uO-nmPThQ1XiEFvQ1GpX6FDqloFMDi96JDJM7WinyDYPxvOVjXkySRu1S3_d4qVe_iyNI6G4w31mJ4RJFxSRi8sj3PugeSGadC0I75L_t7o2vQOnPXDOp1KrymYHWaratdCpNYxKL5xC12jqW6gUGlXc8FvyrflDriClfawjzD6ni38-SLixaDLrtWML5claI28BKJQqu3Gvb_duOWhBnLgwXQQfveMxTbwllqBmpc10SdBOS6Tp-bUjRE0usRpnhnRG_Nhg8KptKQUPg4S6ggNg4jVBp7YZ3xtjg"}
	sc, err := stomp.ConnectOverWS(conn, h)
	if err != nil {
		log.Print("couldn't create stomp connection: %v", err)
		return nil, err
	}

	return sc, nil
}