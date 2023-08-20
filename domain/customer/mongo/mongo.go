package mongo

import (
	"context"
	"time"

	"github.com/bahaa-noah/go-ddd/aggregate"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	db       *mongo.Database
	customer *mongo.Collection
}

type mongoCustomer struct {
	ID   uuid.UUID `bson:"_id"`
	Name string    `bson:"name"`
}

func NewFromCustomer(c aggregate.Customer) mongoCustomer {
	return mongoCustomer{
		ID:   c.GetID(),
		Name: c.GetName(),
	}
}

func (m mongoCustomer) ToAggregate() (aggregate.Customer, error) {
	c := aggregate.Customer{}
	c.SetID(m.ID)
	c.SetName(m.Name)
	return c, nil
}

func New(ctx context.Context, connectionString string) (*MongoRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))

	if err != nil {
		return nil, err
	}

	db := client.Database("ddd")
	customer := db.Collection("customers")

	return &MongoRepository{
		db:       db,
		customer: customer,
	}, nil
}

func (mr *MongoRepository) Get(id uuid.UUID) (aggregate.Customer, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"_id": id}
	var customer mongoCustomer
	err := mr.customer.FindOne(ctx, filter).Decode(&customer)
	if err != nil {
		return aggregate.Customer{}, err
	}
	return customer.ToAggregate()
}

func (mr *MongoRepository) Add(c aggregate.Customer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := mr.customer.InsertOne(ctx, NewFromCustomer(c))

	if err != nil {
		return err
	}

	return nil
}

func (mr *MongoRepository) Update(c aggregate.Customer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": c.GetID()}
	update := bson.M{"$set": bson.M{"name": c.GetName()}}

	_, err := mr.customer.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	return nil
}
