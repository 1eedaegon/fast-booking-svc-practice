package db

import (
	"context"

	"github.com/1eedaegon/fast-booking-svc-practice/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const DBNAME = "reservation"
const USERCOLL = "users"

type UserStore interface {
	GetUserByID(context.Context, id string) (*types.User, error)
}
type PostgresUserStore struct{}
type MongoUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewUserMongoStore(client *mongo.Client) *MongoUserStore {
	return &MongoUserStore{
		client: client,
		coll: client.Database(DBNAME).Collection(USERCOLL)
	}
}

func (s *MongoUserStore) GetUserByID(ctx context.Context, id string) (*types.User, error) {
	var user types.User
	if err := s.coll.FindOne(ctx, bson.M("_id")).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}