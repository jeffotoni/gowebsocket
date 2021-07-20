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
		PORT = "3000"
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
}
