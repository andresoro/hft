package gdax

import (
	"log"
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
	WSConn     *websocket.Conn
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

// Connect connects the client to GDAX websocket
func (c *Client) Connect() {
	var dialer websocket.Dialer

	conn, _, err := dialer.Dial(wsURL, nil)
	if err != nil {
		log.Fatal(err)
	}
	c.WSConn = conn
}

// Subscribe is a method to sub or unsub to a orderbook via a GDAX websocket with a given
// product e.g 'BTC-USD' 'ETH-BTC'
func (c *Client) Subscribe(channelType string, products []string) error {

	subscribe := Message{
		Type: "subscribe",
		Channels: []MessageChannel{
			MessageChannel{
				Name:       channelType,
				ProductIDs: products,
			},
		},
	}

	if err := c.WSConn.WriteJSON(subscribe); err != nil {
		return err
	}

	return nil

}
