package customer

import (
	"errors"

	"github.com/bahaa-noah/go-ddd/aggregate"
	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound       = errors.New("customer not found")
	ErrFailedToAddCustomer    = errors.New("failed to add customer")
	ErrFailedToUpdateCustomer = errors.New("failed to update customer")
)

type CustomerRepository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
}
