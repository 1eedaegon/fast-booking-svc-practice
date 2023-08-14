package types

import (
	"fmt"
	"regexp"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost      = 12
	minFirstNameLen = 2
	minLastNameLen  = 2
	minPasswordLen  = 7
)

type UpdateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (p UpdateUserParams) ToBSON() bson.M {
	m := bson.M{}
	if len(p.FirstName) > 0 {
		m["firstName"] = p.FirstName
	}
	if len(p.LastName) > 0 {
		m["lastName"] = p.LastName
	}
	return m
}

type CreateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json: "email"`
	Password  string `json: "password"`
}

func (params CreateUserParams) Validate() map[string]string {
	errors := map[string]string{}
	if len(params.FirstName) < minFirstNameLen {
		errors["firstName"] = fmt.Sprintf("First name should be at least %d", minFirstNameLen)
	}
	if len(params.LastName) < minLastNameLen {
		errors["lastName"] = fmt.Sprintf("Last name should be at least %d", minLastNameLen)
	}
	if len(params.Password) < minPasswordLen {
		errors["password"] = fmt.Sprintf("Password should be at least %d", minFirstNameLen)
	}
	if !isValidEmail(params.Email) {
		errors["email"] = fmt.Sprintf("Email invalid: %s", params.Email)
	}
	return errors
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(email)
}

type User struct {
	ID                primitive.ObjectID `bson: "_id,omitempty" json:"id,omitempty"`
	FirstName         string             `bson: "firstName" json:"firstName"`
	LastName          string             `bson: "firstName" json:"lastName"`
	Email             string             `bson: "email" json: "email"`
	EncryptedPassword string             `bson: "password" json: "-"`
}

func NewUserFromParams(params CreateUserParams) (*User, error) {
	encPw, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
	if err != nil {
		return nil, err
	}
	return &User{
		FirstName:         params.FirstName,
		LastName:          params.LastName,
		Email:             params.Email,
		EncryptedPassword: string(encPw),
	}, nil
}
