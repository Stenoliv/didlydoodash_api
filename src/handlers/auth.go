package handlers

import (
	"DidlyDoodash-api/src/data"
	"DidlyDoodash-api/src/db"
	"DidlyDoodash-api/src/db/daos"
	"DidlyDoodash-api/src/utils"
	"DidlyDoodash-api/src/utils/jwt"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * Signin Function
 */
type SigninInput struct {
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	RememberMe bool   `json:"rememberMe" default:"false"`
}

func Signin(c *gin.Context) {
	var input SigninInput
	tx := db.DB.Begin()
	// Bind request body
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.InvalidInput)
		return
	}
	fmt.Print(input)
	// Try to get user from database
	user, err := daos.GetUser(input.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.InvalidInput)
		return
	}
	// Validate input.Password with found user password
	ok := user.Validatepassword(input.Password)
	if !ok {
		c.JSON(http.StatusBadRequest, utils.InvalidInput)
		return
	}
	// Generate tokens
	access, err := jwt.GenerateAccessToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ServerError)
		return
	}
	// implement rememberMe
	refresh, err := jwt.GenerateRefreshToken(user.ID, input.RememberMe, tx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ServerError)
		return
	}
	tokens := &utils.Tokens{
		Access:  &access,
		Refresh: &refresh,
	}

	tx.Commit()
	// Send final response
	c.JSON(http.StatusOK, gin.H{"user": user, "tokens": tokens})
}

/**
 * Signup Function
 */
type SignupInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Signup(c *gin.Context) {
	var input SignupInput
	tx := db.DB.Begin()
	// Bind request body
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.InvalidInput)
		return
	}
	// Create new user object
	user := &data.User{Username: input.Username, Email: input.Email, Password: input.Password}
	// Try to save new user to database
	err = user.SaveUser(tx)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, utils.InvalidInput)
		return
	}
	// Generate tokens
	access, err := jwt.GenerateAccessToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ServerError)
		return
	}
	// implement rememberMe
	refresh, err := jwt.GenerateRefreshToken(user.ID, false, tx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ServerError)
		return
	}
	tokens := &utils.Tokens{
		Access:  &access,
		Refresh: &refresh,
	}

	// Commit and send final response
	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"user": user, "tokens": tokens})
}

// Signout function
func Signout(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

// Refresh function
func Refresh(c *gin.Context) {
	// Extract token from request and validate
	tokenStr := jwt.ExtractToken(c)
	token, err := jwt.ValidateToken(tokenStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.AuthenticationError)
		return
	}

	// Extract claims
	claims, err := jwt.ExtractTokenClaims(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.AuthenticationError)
		return
	}

	// Extract subject from token
	sub, err := claims.GetSubject()
	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.AuthenticationError)
		return
	}

	/**
	 * Check for refresh token in db
	 */

	// Check if refresh token
	if claims["type"] != "refresh" {
		c.JSON(http.StatusUnauthorized, utils.AuthenticationError)
		return
	}

	// Generate new access token
	access, err := jwt.GenerateAccessToken(sub)
	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.AuthenticationError)
		return
	}
	tokens := &utils.Tokens{
		Access:  &access,
		Refresh: &tokenStr,
	}

	c.JSON(http.StatusOK, gin.H{"tokens": tokens})
}
