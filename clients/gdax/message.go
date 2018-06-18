package gdax

// SubscribeMessage is a struct that holds data needed to subscribe and unsubscribe
// from GDAX websocket channels
type SubscribeMessage struct {
	Type      string `json:"type"`
	ProductID string `json:"product_id"`
}
