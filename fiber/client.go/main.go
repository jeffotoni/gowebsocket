package main

import (
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/fasthttp/websocket"
)

var addr = flag.String("addr", "localhost:3000", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)
	if len(os.Args) != 2 {
		log.Fatal("Uso: ./client [name]")
	}
	client := os.Args[1]

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/chat/user_" + client}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	err = c.WriteMessage(websocket.TextMessage, []byte(`{"msg":"what's value?"}`))
	if err != nil {
		log.Println("write:", err)
		return
	}

	for {
		select {
		case <-done:
			return
		case <-interrupt:
			log.Println("interrupt")
			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Millisecond):
			}
			return
		}
	}
}
