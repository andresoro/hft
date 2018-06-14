package client

import "net/http"

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
}
