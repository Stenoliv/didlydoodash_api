package handlers

import (
	"DidlyDoodash-api/src/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"user": &data.User{
		Username: "Test",
		Email: "test@gamil.com",
	}})
}

func Signup(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"user": &data.User{
		Username: "Test",
		Email: "test@gamil.com",
	}})
}