package services

import (
	"testing"

	"github.com/bahaa-noah/go-ddd/aggregate"
	"github.com/google/uuid"
)

func products_init(t *testing.T) []aggregate.Product {

	p1, err := aggregate.NewProduct("P1", "Test Product1", 100)
	if err != nil {
		t.Fatalf("Error creating product: %s", err)
	}

	p2, err := aggregate.NewProduct("P2", "Test Product2", 100)
	if err != nil {
		t.Fatalf("Error creating product: %s", err)
	}
	p3, err := aggregate.NewProduct("P3", "Test Product3", 100)
	if err != nil {
		t.Fatalf("Error creating product: %s", err)
	}

	return []aggregate.Product{p1, p2, p3}
}

func TestOrder_NewOrderSvc(t *testing.T) {
	products := products_init(t)

	// Create the service
	oSvc, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)

	if err != nil {
		t.Error(err)
	}

	cust, err := aggregate.NewCustomer("Bahaa")

	if err != nil {
		t.Error(err)
	}

	err = oSvc.customers.Add(cust)

	if err != nil {
		t.Fatalf("Error adding customer: %s", err)
	}

	productIDs := []uuid.UUID{products[0].GetID(), products[1].GetID()}
	// Create a new order
	_, err = oSvc.CreateOrder(cust.GetID(), productIDs)

	if err != nil {
		t.Fatalf("Error creating order: %s", err)
	}

}
