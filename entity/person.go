package entity

import "github.com/google/uuid"

// a Person is an entity the represents a person in all domains
type Person struct {
	ID   uuid.UUID
	Name string
	Age  int
}
