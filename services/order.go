package services

import (
	"context"
	"log"

	"github.com/bahaa-noah/go-ddd/aggregate"
	"github.com/bahaa-noah/go-ddd/domain/customer"
	"github.com/bahaa-noah/go-ddd/domain/customer/memory"
	"github.com/bahaa-noah/go-ddd/domain/customer/mongo"
	"github.com/bahaa-noah/go-ddd/domain/product"
	prodMemory "github.com/bahaa-noah/go-ddd/domain/product/memory"
	"github.com/google/uuid"
)

type OrderConfiguration func(oSvc *OrderService) error

type OrderService struct {
	customers customer.CustomerRepository
	products  product.ProductRepository
}

func NewOrderService(opts ...OrderConfiguration) (*OrderService, error) {
	oSvc := &OrderService{}

	for _, opt := range opts {
		if err := opt(oSvc); err != nil {
			return nil, err
		}
	}

	return oSvc, nil
}

// WithCustomerRepository applies a given customer repository to the OrderService
func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	// return a function that matches the OrderConfiguration alias,
	// You need to return this so that the parent function can take in all the needed parameters
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := memory.New()
	return WithCustomerRepository(cr)
}

func WithMongoCustomerRepository(ctx context.Context, connectionStr string) OrderConfiguration {
	return func(os *OrderService) error {
		mr, err := mongo.New(ctx, connectionStr)
		if err != nil {
			return err
		}
		os.customers = mr
		return nil
	}
}

// WithMemoryProductRepository adds a in memory product repo and adds all input products
func WithMemoryProductRepository(products []aggregate.Product) OrderConfiguration {
	return func(os *OrderService) error {
		// Create the memory repo, if we needed parameters, such as connection strings they could be inputted here
		pr := prodMemory.New()

		// Add Items to repo
		for _, p := range products {
			err := pr.Add(p)
			if err != nil {
				return err
			}
		}
		os.products = pr
		return nil
	}
}

func (oSvc *OrderService) CreateOrder(customerID uuid.UUID, productIDs []uuid.UUID) (float64, error) {
	c, err := oSvc.customers.Get(customerID)
	if err != nil {
		return 0, err
	}

	var products []aggregate.Product
	var totalPrice float64

	for _, id := range productIDs {
		p, err := oSvc.products.GetByID(id)
		if err != nil {
			return 0, err
		}
		products = append(products, p)
		totalPrice += p.GetPrice()
	}

	log.Printf("Customer: %s has ordered %d products", c.GetID(), len(products))

	return totalPrice, nil
}
