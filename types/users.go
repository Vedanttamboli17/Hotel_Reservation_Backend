package types

import (
	"fmt"
	"regexp"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName         string             `bson:"firstName" json:"firstName"`
	LastName          string             `bson:"lastName" json:"lastName"`
	Email             string             `bson:"email" json:"email"`
	EncryptedPassword string             `bson:"encryptedPassword" json:"-"`
}

type UserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UpdateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

const (
	cost            = 10
	MinFirstNameLen = 3
	MinLastNameLen  = 4
	MinPasswordLen  = 7
)

func (params UpdateUserParams) GenerateBSON() bson.M {
	m := bson.M{}
	if len(params.FirstName) > 0 {
		m["firstName"] = params.FirstName
	}
	if len(params.LastName) > 0 {
		m["lastName"] = params.LastName
	}
	return m
}

func (params UserParams) Validate() map[string]string {
	err := map[string]string{}
	if len(params.FirstName) < MinFirstNameLen {
		err["firstName"] = fmt.Sprintf("The minimum length of firstName should be %d", MinFirstNameLen)
	}
	if len(params.LastName) < MinLastNameLen {
		err["lastName"] = fmt.Sprintf("The minimum length of lastName should be %d", MinLastNameLen)
	}
	if len(params.Password) < MinPasswordLen {
		err["email"] = fmt.Sprintf("The minimum length of password should be %d", MinPasswordLen)
	}
	if !isEmailValid(params.Email) {
		err["password"] = "The email is invalid"
	}
	return err
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(e)
}

func NewUserFromParams(params UserParams) (*User, error) {
	encryptPass, err := bcrypt.GenerateFromPassword([]byte(params.Password), cost)
	if err != nil {
		return nil, err
	}
	return &User{
		FirstName:         params.FirstName,
		LastName:          params.LastName,
		Email:             params.Email,
		EncryptedPassword: string(encryptPass),
	}, nil
}
