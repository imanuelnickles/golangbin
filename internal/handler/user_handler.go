package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type User interface {
	Get(c *gin.Context)
}

type userHandler struct{}

func NewUserHandler() User {
	return userHandler{}
}

func (h userHandler) Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": time.Now().String(),
		"user_id": 123,
	})
}
