package users

import (
	"net/http"
	"user-identity-subdomain/domain/users"
	"user-identity-subdomain/rest_errors"
	"user-identity-subdomain/services"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	var request users.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	if request.Email == "" {
		restErr := rest_errors.NewBadRequestError("email should not be empty")
		c.JSON(restErr.Status(), restErr)
		return
	}

	if request.Password == "" {
		restErr := rest_errors.NewBadRequestError("Password should be empty")
		c.JSON(restErr.Status(), restErr)
		return
	}

	user, err := services.UsersService.LoginUser(request)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	var tokenRes users.LoginResponse

	tokenRes.AccessToken = user.AccessToken
	tokenRes.Email = user.Email

	c.JSON(http.StatusOK, tokenRes)

}
