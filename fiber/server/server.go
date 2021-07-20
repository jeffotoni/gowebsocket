/*
* Golang presentation
*
* @package     main
* @author      @jeffotoni
* @size        2017
 */

package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/websocket"
)

var (
	HOST = os.Getenv("HOST")
	PORT = os.Getenv("PORT")
)

func init() {
	if len(PORT) == 0 {
		PORT = "1234"
	}
	if len(HOST) == 0 {
		HOST = "0.0.0.0"
	}
}

func ListenWebSocker(ws *websocket.Conn) {
	for { // loop
		var reply string //  Receive the message by reference
		// Receiving client message
		websocket.Message.Receive(ws, &reply)
		if len(reply) > 0 {
			msg := `{"name":"jeffotoni", "code":"kong engdb => ` + reply + `"}`
			websocket.Message.Send(ws, msg) // Sending message to the client
			time.Sleep(time.Millisecond * 600)
		}
	}
}

func main() {
	http.Handle("/chat", websocket.Handler(ListenWebSocker))
	log.Println("\u001B[1;33mRun Server Websocket=>"+HOST+":"+PORT, "\033[0m")
	if err := http.ListenAndServe(HOST+":"+PORT, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
