package internal

import "time"

type Event interface{}

type ProductShipped struct {
	SKU      string
	Quantity int
	DateTime time.Time
}

type ProductReceived struct {
	SKU      string
	Quantity int
	DateTime time.Time
}

type InventoryAdjusted struct {
	SKU      string
	Quantity int
	Reason   string
	DateTime time.Time
}
