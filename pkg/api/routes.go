package api

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	router            = gin.Default()
	ErrInvalidDetails = errors.New("you have entered invalid details")
)

func Run() {
	router.POST("/login", Login)
	log.Fatal(router.Run(":8080"))
}
