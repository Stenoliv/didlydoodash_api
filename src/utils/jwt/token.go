package jwt

import (
	"DidlyDoodash-api/src/data"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

// Generate a new access token
func GenerateAccessToken(userID data.Nanoid) (string, error) {
	lifespan := GetAccessTime()
	jti, err := gonanoid.New()
	if err != nil {
		return "", err
	}
	claims := jwt.MapClaims{}
	claims["jti"] = jti
	claims["sub"] = userID
	claims["user"] = data.CurrentUser
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(lifespan)).Unix()
	claims["type"] = "access"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
}

// Generate a new refresh token
func GenerateRefreshToken(userID data.Nanoid, jti data.Nanoid) (string, error) {
	lifespan := GetRefreshTime()
	claims := jwt.MapClaims{}
	claims["jti"] = jti
	claims["sub"] = userID
	claims["user"] = data.CurrentUser
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(lifespan)).Unix()
	claims["type"] = "access"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
}

// Exctract access token from cookie or Authorization header
func ExtractAccessToken(c *gin.Context) string {
	if access_token, err := c.Request.Cookie("access_token"); err == nil {
		return access_token.Value
	}
	if bearerToken := c.Request.Header.Get("Authorization"); bearerToken != "" && len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	if queryToken := c.Query("access_token"); queryToken != "" {
		return queryToken
	}
	return ""
}

// Exctract refresh token from cookie or Authorization header
func ExtractRefreshToken(c *gin.Context) string {
	if refresh_token, err := c.Request.Cookie("refresh_token"); err == nil {
		return refresh_token.Value
	}
	if bearerToken := c.Request.Header.Get("Authorization"); bearerToken != "" && len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	if queryToken := c.Query("refresh_token"); queryToken != "" {
		return queryToken
	}
	return ""
}

// Extract payload from token
func ExtractTokenClaims(tokenIn string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenIn, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}

// Get access token time from .env
func GetAccessTime() int {
	time, err := strconv.Atoi(os.Getenv("TOKEN_LIFESPAN_ACCESS"))
	if err != nil {
		return 0
	}
	return time * 60
}

// Get refresh token time from .env
func GetRefreshTime() int {
	time, err := strconv.Atoi(os.Getenv("TOKEN_LIFESPAN_REFRESH"))
	if err != nil {
		return 0
	}
	return time * 60
}
