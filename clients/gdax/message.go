package gdax

import "time"

// Message is a struct that holds data needed send and recieve data
// from GDAX websocket channels
type Message struct {
	Type          string           `json:"type"`
	ProductID     string           `json:"product_id"`
	ProductIDs    []string         `json:"product_ids"`
	TradeID       int              `json:"trade_id,number"`
	OrderID       string           `json:"order_id"`
	Sequence      int64            `json:"sequence,number"`
	MakerOrderID  string           `json:"maker_order_id"`
	TakerOrderID  string           `json:"taker_order_id"`
	Time          time.Time        `json:"time,string"`
	RemainingSize float64          `json:"remaining_size,string"`
	NewSize       float64          `json:"new_size,string"`
	OldSize       float64          `json:"old_size,string"`
	Size          float64          `json:"size,string"`
	Price         float64          `json:"price,string"`
	Side          string           `json:"side"`
	Reason        string           `json:"reason"`
	OrderType     string           `json:"order_type"`
	Funds         float64          `json:"funds,string"`
	NewFunds      float64          `json:"new_funds,string"`
	OldFunds      float64          `json:"old_funds,string"`
	Message       string           `json:"message"`
	Bids          [][]string       `json:"bids,omitempty"`
	Asks          [][]string       `json:"asks,omitempty"`
	Changes       [][]string       `json:"changes,omitempty"`
	LastSize      float64          `json:"last_size,string"`
	BestBid       float64          `json:"best_bid,string"`
	BestAsk       float64          `json:"best_ask,string"`
	Channels      []MessageChannel `json:"channels"`
	UserID        string           `json:"user_id"`
	ProfileID     string           `json:"profile_id"`
}

// MessageChannel is which channels to subscribe to
type MessageChannel struct {
	Name       string   `json:"name"`
	ProductIDs []string `json:"product_ids"`
}
