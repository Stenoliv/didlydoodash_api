package handlers

import (
	"DidlyDoodash-api/src/db"
	"DidlyDoodash-api/src/db/daos"
	"DidlyDoodash-api/src/db/models"
	"DidlyDoodash-api/src/utils"
	"DidlyDoodash-api/src/utils/jwt"
	"net/http"
	"time"

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
		tx.Rollback()
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
	user := &models.User{Username: input.Username, Email: input.Email, Password: input.Password}
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
		tx.Rollback()
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
	unauthorizedResponse := func() {
		c.JSON(http.StatusUnauthorized, utils.AuthenticationError)
	}

	// Extract token from request without validation
	tokenStr := jwt.ExtractToken(c)
	token, err := jwt.ParseTokenWithoutValidation(tokenStr) // Parse without validation
	if err != nil {
		utils.LogError(err, "Failed to parse token")
		unauthorizedResponse()
		return
	}

	claims, err := jwt.ExtractTokenClaims(token)
	if err != nil {
		utils.LogError(err, "Failed to extract claims")
		unauthorizedResponse()
		return
	}

	// Extract subject from token
	sub, err := claims.GetSubject()
	if err != nil {
		utils.LogError(err, "Failed to extract subject")
		unauthorizedResponse()
		return
	}

	// Check if it's a refresh token
	if claims["type"] != "refresh" {
		utils.LogError(err, "Token not refresh")
		unauthorizedResponse()
		return
	}

	// Check for token in database
	var session models.UserSession
	jti, ok := claims["jti"].(string)
	if !ok {
		utils.LogError(err, "Failed to retrieve jti")
		unauthorizedResponse()
		return
	}

	if session, err = daos.GetSession(sub, jti); err != nil {
		utils.LogError(err, "Didn't find a session")
		unauthorizedResponse()
		return
	}

	// Now check if the session is still valid
	if session.ExpireDate.After(time.Now()) {
		_, err = jwt.ValidateToken(tokenStr)
		if err != nil {
			utils.LogError(err, "Not a valid token")
			unauthorizedResponse()
			return
		}

		// Generate new access token
		access, err := jwt.GenerateAccessToken(sub)
		if err != nil {
			utils.LogError(err, "Failed to generate a access token")
			unauthorizedResponse()
			return
		}
		tokens := &utils.Tokens{
			Access:  &access,
			Refresh: &tokenStr,
		}

		c.JSON(http.StatusOK, gin.H{"tokens": tokens})
	} else {
		// Session has expired in the database
		if err = db.DB.Delete(&session, "user_id = ? AND jti = ?", sub, jti).Error; err != nil {
			utils.LogError(err, "Failed to delete old record in database")
			unauthorizedResponse()
			return
		}
		unauthorizedResponse()
	}
}
