package aggregate

import (
	"errors"

	"github.com/bahaa-noah/go-ddd/entity"
	"github.com/google/uuid"
)

var (
	ErrMissingValue = errors.New("missing value")
)

type Product struct {
	item     *entity.Item
	price    float64
	quantity int
}

func NewProduct(name string, description string, price float64) (Product, error) {
	if name == "" || description == "" {
		return Product{}, ErrMissingValue
	}

	return Product{
		item: &entity.Item{
			ID:          uuid.New(),
			Name:        name,
			Description: description,
		},
		price:    price,
		quantity: 0,
	}, nil
}

func (p Product) GetID() uuid.UUID {
	return p.item.ID
}

func (p Product) GetName() string {
	return p.item.Name
}

func (p Product) GetDescription() string {
	return p.item.Description
}

func (p Product) GetPrice() float64 {
	return p.price
}

func (p Product) GetQuantity() int {
	return p.quantity
}

func (p Product) GetItem() *entity.Item {
	return p.item
}

func (p *Product) SetQuantity(quantity int) {
	p.quantity = quantity
}
