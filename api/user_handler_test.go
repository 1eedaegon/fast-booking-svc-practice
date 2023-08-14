package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/1eedaegon/fast-booking-svc-practice/db"
	"github.com/1eedaegon/fast-booking-svc-practice/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	testdburi  = "mongodb://localhost:27017"
	testdbname = "test-reservation"
)

type testdb struct {
	db.UserStore
}

func (tdb testdb) teardown(t *testing.T) {

	if err := tdb.UserStore.Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}

func setup(t *testing.T) *testdb {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(testdburi))
	if err != nil {
		t.Fatal(err)
	}
	return &testdb{
		UserStore: db.NewUserMongoStore(client, testdbname),
	}
}
func TestPostUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)

	app := fiber.New()
	userHandler := NewUserHandler(tdb.UserStore)
	app.Post("/", userHandler.HandlePostUser)

	params := types.CreateUserParams{
		Email:     "some@foo.com",
		FirstName: "Lorem",
		LastName:  "ipsum",
		Password:  "12345678910",
	}
	b, _ := json.Marshal(params)
	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp.Status)

	// 1. ioutil에서 한번에 읽는 방법
	// bb, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	t.Error(err)
	// }
	// fmt.Println(string(bb))

	// 2. Decode 시켜서 읽는 방법
	var user types.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		t.Error(err)
	}
	if user.Email != params.Email {
		t.Errorf("Expected email: %s but got: %s", params.Email, user.Email)
	}
	if user.FirstName != params.FirstName {
		t.Errorf("Expected first name: %s but got: %s", params.FirstName, user.FirstName)
	}
	if user.LastName != params.LastName {
		t.Errorf("Expected last name: %s but got: %s", params.LastName, user.LastName)
	}
	fmt.Println(user)
}
