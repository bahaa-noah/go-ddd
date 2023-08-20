package memory

import (
	"fmt"
	"sync"

	"github.com/bahaa-noah/go-ddd/aggregate"
	"github.com/bahaa-noah/go-ddd/domain/product"
	"github.com/google/uuid"
)

type MemoryProductRepository struct {
	products map[uuid.UUID]aggregate.Product
	sync.Mutex
}

func New() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[uuid.UUID]aggregate.Product),
	}
}

func (mpr *MemoryProductRepository) GetAll() ([]aggregate.Product, error) {
	var products []aggregate.Product
	for _, product := range mpr.products {
		products = append(products, product)
	}
	return products, nil
}

func (mpr *MemoryProductRepository) GetByID(id uuid.UUID) (aggregate.Product, error) {
	if product, ok := mpr.products[id]; ok {
		return product, nil
	}
	return aggregate.Product{}, product.ErrProductNotFound
}

func (mpr *MemoryProductRepository) Add(newProd aggregate.Product) error {

	if mpr.products == nil {
		mpr.Lock()
		defer mpr.Unlock()
		mpr.products = make(map[uuid.UUID]aggregate.Product)
	}

	if _, ok := mpr.products[newProd.GetID()]; ok {
		return fmt.Errorf("%w: %s", product.ErrProductAlreadyExists, newProd.GetName())
	}

	mpr.Lock()
	defer mpr.Unlock()
	mpr.products[newProd.GetID()] = newProd
	return nil
}

func (mpr *MemoryProductRepository) Update(product aggregate.Product) error {
	mpr.Lock()
	defer mpr.Unlock()
	mpr.products[product.GetID()] = product
	return nil
}

func (mpr *MemoryProductRepository) Delete(id uuid.UUID) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[id]; !ok {
		return product.ErrProductNotFound
	}

	delete(mpr.products, id)
	return nil
}
