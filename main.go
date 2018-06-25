package main

import (
	"fmt"

	"github.com/andresoro/hft/clients/gdax"
)

func main() {
	// init client
	client := gdax.NewClient("test", "test", "test")
	products := []string{"ETH-USD"}

	// connect to WS with desired product
	client.Connect()
	client.Subscribe("level2", products)

	// infinite for loop which reads incoming messages
	var msg gdax.Message
	for true {

		if err := client.WSConn.ReadJSON(&msg); err != nil {
			fmt.Println(err.Error())
			break
		}
		fmt.Println(msg.Side, msg.Price)

	}
}
