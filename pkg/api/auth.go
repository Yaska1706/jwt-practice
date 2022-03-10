package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yaska1706/jwt-practice/internal/auth"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var user1 = User{
	Username: "username1",
	Password: "pass1",
}

func Login(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid")
		return
	}
	if err := IsUserPresent(u); err != nil {
		c.JSON(http.StatusUnauthorized, ErrInvalidDetails)
	}

	token, err := auth.CreateToken(u.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	tokens := map[string]string{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
	}

	c.JSON(http.StatusOK, tokens)

}

func IsUserPresent(u User) error {

	if user1.Username != u.Username || user1.Password != u.Password {

		return ErrInvalidDetails
	}
	return nil
}
