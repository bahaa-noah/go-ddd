// In memory implementation of customer repository
package memory

import (
	"fmt"
	"sync"

	"github.com/bahaa-noah/go-ddd/aggregate"
	"github.com/bahaa-noah/go-ddd/domain/customer"
	"github.com/google/uuid"
)

type MemoryRepository struct {
	customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]aggregate.Customer),
	}
}

func (mr *MemoryRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	if customer, ok := mr.customers[id]; ok {
		return customer, nil
	}
	return aggregate.Customer{}, customer.ErrCustomerNotFound
}

func (mr *MemoryRepository) Add(c aggregate.Customer) error {
	if mr.customers == nil {
		mr.Lock()
		defer mr.Unlock()
		mr.customers = make(map[uuid.UUID]aggregate.Customer)
	}

	//make sure the customer is not in the repository
	if _, ok := mr.customers[c.GetID()]; ok {
		return fmt.Errorf("customer already exists :%w", customer.ErrFailedToAddCustomer)
	}

	mr.Lock()
	defer mr.Unlock()
	mr.customers[c.GetID()] = c
	return nil
}

func (mr *MemoryRepository) Update(c aggregate.Customer) error {
	if _, ok := mr.customers[c.GetID()]; !ok {
		return fmt.Errorf("customer doesn't exists :%w", customer.ErrFailedToUpdateCustomer)
	}

	mr.Lock()
	defer mr.Unlock()
	mr.customers[c.GetID()] = c
	return nil
}
