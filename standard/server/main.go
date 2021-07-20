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
	var reply string
	for {
		if err := websocket.Message.Receive(ws, &reply); err != nil {
			// handle error
			if err.Error() == "EOF" {
				break
			}
			log.Println("error Receive:", err)
			continue
		}

		// send message
		msg := `{"name":"jeffotoni", "code":"kong server engdb => ` + reply + `"}`
		if err := websocket.Message.Send(ws, msg); err != nil {
			// handle error
			log.Println("error Send:", err)
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
