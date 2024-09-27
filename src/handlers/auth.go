package handlers

import (
	"DidlyDoodash-api/src/data"
	"DidlyDoodash-api/src/db"
	"DidlyDoodash-api/src/db/daos"
	"DidlyDoodash-api/src/utils"
	"DidlyDoodash-api/src/utils/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SigninInput struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	RememberMe bool   `json:"rememberMe"`
}

func Signin(c *gin.Context) {
	var input SigninInput
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.InvalidInput )
		return
	}
	user, err := daos.GetUser(input.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ServerError)
		return
	}
	if user == (data.User{}) {
		c.JSON(http.StatusBadRequest, utils.InvalidInput)
		return
	} 
	ok := user.Validatepassword(input.Password)
	if !ok {
		c.JSON(http.StatusBadRequest, utils.InvalidInput)
		return
	}
	token, err := jwt.GenerateAccessToken(data.Nanoid(user.ID)) 
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ServerError)
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"user": user, "token": token})
}

type SignupInput struct {
	Username string `json:"username" binding:"required"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

func Signup(c *gin.Context) {
	var input SignupInput
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.InvalidInput)
		return
	}
	user := &data.User{Username: input.Username, Email: input.Email, Password: input.Password}
	tx := db.DB.Begin()
	err = user.SaveUser(tx)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, utils.InvalidInput)
		return
	} 
	token, err := jwt.GenerateAccessToken(data.Nanoid(user.ID)) 
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, utils.ServerError)
		return
	}
	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"user": user, "token": token})
}

func Signout(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func Refresh(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
