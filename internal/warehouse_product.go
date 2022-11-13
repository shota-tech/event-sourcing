package internal

import (
	"errors"
	"time"
)

type (
	WarehouseProduct struct {
		SKU          string
		events       []Event
		currentState CurrentState
	}

	CurrentState struct {
		QuantityOnHand int
	}
)

func NewWarehouseProduct(sku string) *WarehouseProduct {
	return &WarehouseProduct{
		SKU:    sku,
		events: []Event{},
	}
}

func (p *WarehouseProduct) Events() []Event {
	return p.events
}

func (p *WarehouseProduct) QuantityOnHand() int {
	return p.currentState.QuantityOnHand
}

func (p *WarehouseProduct) ShipProduct(quantity int) error {
	if quantity > p.currentState.QuantityOnHand {
		return errors.New("not enough product to ship")
	}
	event := ProductShipped{
		SKU:      p.SKU,
		Quantity: quantity,
		DateTime: time.Now(),
	}
	p.AddEvent(event)
	return nil
}

func (p *WarehouseProduct) ReceiveProduct(quantity int) {
	event := ProductReceived{
		SKU:      p.SKU,
		Quantity: quantity,
		DateTime: time.Now(),
	}
	p.AddEvent(event)
}

func (p *WarehouseProduct) AdjustInventory(quantity int, reason string) error {
	if p.currentState.QuantityOnHand+quantity < 0 {
		return errors.New("cannot adjust to a negative quantity on hand")
	}
	event := InventoryAdjusted{
		SKU:      p.SKU,
		Quantity: quantity,
		Reason:   reason,
		DateTime: time.Now(),
	}
	p.AddEvent(event)
	return nil
}

func (p *WarehouseProduct) AddEvent(event Event) error {
	switch e := event.(type) {
	case ProductShipped:
		p.applyProductShipped(e)
	case ProductReceived:
		p.applyProductReceived(e)
	case InventoryAdjusted:
		p.applyInventoryAdjusted(e)
	default:
		return errors.New("unsupported event")
	}
	p.events = append(p.events, event)
	return nil
}

func (p *WarehouseProduct) applyProductShipped(event ProductShipped) {
	p.currentState.QuantityOnHand -= event.Quantity
}

func (p *WarehouseProduct) applyProductReceived(event ProductReceived) {
	p.currentState.QuantityOnHand += event.Quantity

}

func (p *WarehouseProduct) applyInventoryAdjusted(event InventoryAdjusted) {
	p.currentState.QuantityOnHand += event.Quantity
}
