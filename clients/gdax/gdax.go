package gdax

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	baseURL = "https://api.gdax.com"
	wsURL   = "wss://ws-feed.gdax.com"
)

// Client is the client for accessing GDAX websocket feed and
// REST API to read data and make trades.
type Client struct {
	Secret     string
	Key        string
	Pass       string
	HTTPClient *http.Client
	WSConn     websocket.Conn
}

// NewClient returns a GDAX client
func NewClient(secret, key, pass string) *Client {
	client := Client{
		Secret: secret,
		Key:    key,
		Pass:   pass,
	}

	return &client
}

// Subscribe is a method to sub or unsub to a orderbook via a GDAX websocket with a given
// product e.g 'BTC-USD' 'ETH-BTC'
func (c *Client) Subscribe(product string, unsubscribe bool) error {
	var message SubscribeMessage

	if unsubscribe {
		message = SubscribeMessage{"unsubscribe", product}
	} else {
		message = SubscribeMessage{"subscribe", product}
	}

	json, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = c.WSConn.WriteMessage(websocket.TextMessage, json)
	if err != nil {
		return err
	}

	return nil
}
