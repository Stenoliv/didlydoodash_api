package handlers

import (
	"DidlyDoodash-api/src/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SigninType struct  {
	Email string `json:"email"`
	Password string `json:"password"`
}

func Signin(c *gin.Context) {
	var input SigninType
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, "bad")
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": &data.User{
		Username: "Logged in",
		Email: input.Email,
	}})
}


type SignupType struct  {
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func Signup(c *gin.Context) {
	var input SignupType
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, "bad")
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": &data.User{
		Username: input.Username,
		Email: input.Email,
	}})
}