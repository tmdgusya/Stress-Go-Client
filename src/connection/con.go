package connection

import (
	"log"
	"net"

	"github.com/gmallard/stompngo"
)

const SERVER_DOMAIN = "127.17.0.1:57304"

// {
//     "url": "ws://localhost:8761/stomp",
//     "headers": {
//         "token": "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJpc3MiOiJjbGVhbmluZ19sYWIiLCJleHAiOjE2MzI3MjkwNTksInRva2VuX2RhdGEiOnsiYXBwSWQiOjAsImFsbG93T3JpZ2lucyI6bnVsbH19.t-WlN060FqpVQCzBxiCdmur24lh-oXcPNK_q5IJKXRzec3-SdCDXRU6S75j39Q56EEn5YktBaSsjbP4MLynjS8EplfZdjpUkSodjQtZsJxN-hhO_rbU_5noWz5oaZGE_Ou2qGpH038fCIwytb6kqc8nF4sl8xlVCZfenI75kOxc1P8Y_vQAMwQxyYn26-dwlnXV4vKkEK9R5fASCINFAnjzYOtxtAOGG-QBntHY6HVn3U89vDW3ccHZoAEQ2whj9P0z2jbFvc7if-lfBnw-SWhqacq5QMRcQ1yhSwr6kf5bX4oCU0ouDl5MSHOl4Zq3U3fpD1rTpKBdHbBteUiuHMg",
//         "host": "development-app-0"
//     }
// }

func ConnectFactory() (net.Conn, error) {
	netconn, err := net.Dial("tcp", "localhost:61613")
    
	if err != nil {
		log.Println("1 : ", err)
		return nil, nil;
	}

    h := stompngo.Headers{stompngo.HK_ACCEPT_VERSION, "1.1",
        stompngo.HK_HOST, "development-app-0"}
	
	conn, err := stompngo.Connect(netconn, h)

	log.Println(conn);

	if err != nil {
		log.Println("2 : ", err)
		return nil, nil;
	}

    subHead := stompngo.Headers{stompngo.HK_DESTINATION, "/topic/chat"}


    sub, err := conn.Subscribe(subHead)

	if err != nil {
		log.Println("3 : ", err)
		return nil, nil;
	}

    for {
        msg := <-sub

        log.Println("Message : ", msg.Message.BodyString())

    }
}