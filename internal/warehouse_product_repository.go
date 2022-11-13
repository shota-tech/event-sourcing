package internal

type WarehouseProductRepository struct {
	inMemoryStreams map[string][]Event
}

func NewWarehouseProductRepository() *WarehouseProductRepository {
	return &WarehouseProductRepository{
		inMemoryStreams: make(map[string][]Event),
	}
}

func (r *WarehouseProductRepository) Get(sku string) *WarehouseProduct {
	warehouseProduct := NewWarehouseProduct(sku)
	for _, event := range r.inMemoryStreams[sku] {
		warehouseProduct.AddEvent(event)
	}
	return warehouseProduct
}

func (r *WarehouseProductRepository) Save(warehouseProduct *WarehouseProduct) {
	if warehouseProduct == nil {
		return
	}
	r.inMemoryStreams[warehouseProduct.SKU] = warehouseProduct.Events()
}
