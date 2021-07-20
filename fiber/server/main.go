package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
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

func main() {
	app := fiber.New()
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/chat/:id", websocket.New(func(c *websocket.Conn) {
		// c.Locals is added to the *websocket.Conn
		log.Println("allowed:", c.Locals("allowed"))  // true
		log.Println("id:", c.Params("id"))            // 123
		log.Println("query:v=>", c.Query("v"))        // 1.0
		log.Println("session:", c.Cookies("session")) // ""

		for {
			log.Println("new client:", c.Params("id"))
			mt, msgCli, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			jsonStr := `{"msg":"test kong.engdb => "` + string(msgCli) + `","value":1398390.57}`
			err = c.WriteMessage(mt, []byte(jsonStr))
			if err != nil {
				log.Println("cliente saiu:", c.Params("id"))
				continue
			}
		}
	}))
	log.Fatal(app.Listen(ADDR))
}
