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

	"golang.org/x/net/websocket"
)

var (
	HOST = os.Getenv("HOST")
	PORT = os.Getenv("PORT")
	ADDR = ""
)

func init() {
	if len(PORT) == 0 {
		PORT = "3000"
	}
	if len(HOST) == 0 {
		HOST = "0.0.0.0"
	}
	ADDR = HOST + ":" + PORT
}

func ListenWebSocker(ws *websocket.Conn) {
	for { // loop
		var reply string //  Receive the message by reference
		// Receiving client message
		websocket.Message.Receive(ws, &reply)
		if len(reply) > 0 {
			msg := `{"name":"jeffotoni", "code":"kong engdb => ` + reply + `"}`
			websocket.Message.Send(ws, msg) // Sending message to the client
			//time.Sleep(time.Millisecond * 600)
		}
	}
}

func main() {
	http.Handle("/chat", websocket.Handler(ListenWebSocker))
	log.Println("\u001B[1;33mRun Server Websocket=>"+HOST+":"+PORT, "\033[0m")
	if err := http.ListenAndServe(ADDR, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
