package memory

import (
	"testing"

	"github.com/bahaa-noah/go-ddd/aggregate"
	"github.com/bahaa-noah/go-ddd/domain/customer"
	"github.com/google/uuid"
)

func TestMemory_GetCustomer(t *testing.T) {
	type testCase struct {
		test        string
		id          uuid.UUID
		expectedErr error
	}

	cust, err := aggregate.NewCustomer("Bahaa")
	if err != nil {
		t.Fatalf("Failed to create customer: %v", err)
	}

	id := cust.GetID()

	repo := MemoryRepository{
		customers: map[uuid.UUID]aggregate.Customer{
			id: cust,
		},
	}

	testCases := []testCase{
		{
			test:        "no customer by id",
			id:          uuid.MustParse("00000000-0000-0000-0000-000000000000"),
			expectedErr: customer.ErrCustomerNotFound,
		}, {
			test:        "customer by id",
			id:          id,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := repo.Get(tc.id)
			if err != tc.expectedErr {
				t.Errorf("Expected error: %v, got: %v", tc.expectedErr, err)
			}
		})
	}

}
