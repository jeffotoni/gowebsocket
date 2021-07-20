package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/net/websocket"
)

var (
	PORT   = os.Getenv("PORT")
	HOST   = os.Getenv("HOST")
	URL    = os.Getenv("URL")
	URI    = ""
	ORIGIN = "http://localhost/"
)

func init() {
	if len(PORT) == 0 {
		PORT = "1234"
	}
	if len(HOST) == 0 {
		HOST = "ws://localhost"
	}
	if len(URL) == 0 {
		URL = "/chat"
	}

	URI = HOST + ":" + PORT + URL
}

func main() {
	for {
		ws, err := websocket.Dial(URI, "", ORIGIN)
		if err != nil {
			log.Fatal(err)
		}

		if _, err := ws.Write([]byte("kong-client-one")); err != nil {
			log.Fatal(err)
		}

		// for {
		var msg = make([]byte, 512)
		var n int
		if n, err = ws.Read(msg); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Received: %s.\n", msg[:n])
		time.Sleep(time.Millisecond * 2000)
	}

	// // create connection
	// // schema can be ws:// or wss://
	// // host, port – WebSocket server
	// conn, err := websocket.Dial("ws://localhost:1234/chat", "", "http://localhost/")
	// if err != nil {
	//   // handle error
	//   log.Println(err)
	//   return
	// }
	// defer conn.Close()

	// // send message
	// if err = websocket.JSON.Send(conn, `{"msg":"jeffotoni", "code":"xxxxxxxxxxxxe39393"}`); err != nil {
	//   // handle error
	// }

	// // receive message
	// // messageType initializes some type of message
	// message := messageType{}
	// if err := websocket.JSON.Receive(conn, &message); err != nil {
	//   // handle error
	//   log.Println(err)
	// }
}
