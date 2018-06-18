package gdax

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	baseURL = "https://api.gdax.com"
	wsURL   = "wss://ws-feed.gdax.com"
)

// GDAXClient is the client for accessing GDAX websocket feed and
// REST API to read data and make trades.
type GDAXClient struct {
	Secret     string
	Key        string
	Pass       string
	HTTPClient *http.Client
	WSConn     websocket.Conn
}

// NewClient returns a GDAX client
func NewClient(secret, key, pass string) *GDAXClient {
	client := GDAXClient{
		Secret: secret,
		Key:    key,
		Pass:   pass,
	}

	return &client
}
