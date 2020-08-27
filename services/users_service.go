package services

import (
	"fmt"
	"log"
	"time"
	"user-identity-subdomain/domain/users"
	"user-identity-subdomain/rest_errors"

	"github.com/dgrijalva/jwt-go"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct{}

type usersServiceInterface interface {
	LoginUser(users.LoginRequest) (*users.User, rest_errors.RestErr)
}

func (s *usersService) LoginUser(request users.LoginRequest) (*users.User, rest_errors.RestErr) {

	dao := &users.User{}

	dao.Email = request.Email
	dao.Password = request.Password

	err := dao.FindByEmailAndPassword()

	if err != nil {
		restErrDb1 := rest_errors.NewNotFoundError("Database Error")
		return nil, restErrDb1
	}

	if dao.Password == request.Password {
		if dao.TokenStatus == "enabled" {

			tokenString, timeStamp := GenerateToken(dao.Email)

			dao.AccessToken = tokenString
			dao.DateCreated = timeStamp

			fmt.Println("dao.AccessToken:", dao.AccessToken)
			fmt.Println("dao.DateCreated:", dao.DateCreated)

			err = dao.UpdateToken()

		}
	} else {
		restErrDb := rest_errors.NewUnauthorizedError("Invalid User")
		return nil, restErrDb
	}

	return dao, nil
}

func GenerateToken(inputEmail string) (string, string) {

	secret := "mysecrets"

	now := time.Now()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": inputEmail,
		"iss":   "feedback",
	})

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		log.Fatal(err)
	}

	return tokenString, now.String()

}
