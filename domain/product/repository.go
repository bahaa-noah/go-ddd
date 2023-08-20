package product

import (
	"errors"

	"github.com/bahaa-noah/go-ddd/aggregate"
	"github.com/google/uuid"
)

var (
	ErrProductNotFound       = errors.New("product not found")
	ErrFailedToAddProduct    = errors.New("failed to add product")
	ErrFailedToUpdateProduct = errors.New("failed to update product")
	ErrFailedToDeleteProduct = errors.New("failed to delete product")
	ErrProductAlreadyExists  = errors.New("product already exists")
)

type ProductRepository interface {
	GetAll() ([]aggregate.Product, error)
	GetByID(id uuid.UUID) (aggregate.Product, error)
	Add(product aggregate.Product) error
	Update(product aggregate.Product) error
	Delete(id uuid.UUID) error
}
