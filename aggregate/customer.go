// Package aggregate holds aggregates that combines many entities into a full object
package aggregate

import (
	"errors"

	"github.com/bahaa-noah/go-ddd/entity"
	"github.com/bahaa-noah/go-ddd/valueobject"
	"github.com/google/uuid"
)

var (
	ErrInvalidPersonName = errors.New("a customer has to have a valid person name")
)

type Customer struct {
	//Person is the root entity of the Customer
	// Which means Person.ID is the main identifier of the Customer
	person   *entity.Person
	products []*entity.Item

	transactions []valueobject.Transaction
}

func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPersonName
	}
	person := &entity.Person{
		Name: name,
		ID:   uuid.New(),
	}

	return Customer{
		person:       person,
		products:     make([]*entity.Item, 0),
		transactions: make([]valueobject.Transaction, 0),
	}, nil
}

func (c *Customer) GetID() uuid.UUID {
	return c.person.ID
}
func (c *Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.ID = id
}

func (c *Customer) GetName() string {
	return c.person.Name
}
func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.Name = name
}
