package main

import (
	"context"
	"flag"
	"log"

	"github.com/1eedaegon/fast-booking-svc-practice/api"
	"github.com/1eedaegon/fast-booking-svc-practice/db"
	"github.com/1eedaegon/fast-booking-svc-practice/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://localhost:27017"
const dbname = "reservation"
const userColl = "users"

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	coll := client.Database(dbname).Collection(userColl)
	user := types.User{FirstName: "lee", LastName: "daegon"}
	res, err := coll.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}
	var gon types.User
	coll.FindOne(ctx, bson.M{}).Decode(&gon)

	port := flag.String("port", ":8888", "Port of ")
	flag.Parse()

	// TODO: Impl handler
	userHandler := api.NewUserHandler(db.NewUserMongoStore(client))

	app := fiber.New()
	appV1 := app.Group("/app/v1")
	appV1.Get("/user", userHandler.HandleGetUsers)
	appV1.Get("/user/:id", userHandler.HandleGetUser)
	app.Listen(*port)

}
