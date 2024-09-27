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
	"gorm.io/gorm"
)

// Generate a new access token
func GenerateAccessToken(userID string) (string, error) {
	lifespan, err := GetAccessTime()
	if err != nil {
		return "", err
	}
	jti, err := gonanoid.New()
	if err != nil {
		return "", err
	}
	claims := jwt.MapClaims{}
	claims["jti"] = jti
	claims["sub"] = userID
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(lifespan)).Unix()
	claims["type"] = "access"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
}

// Generate a new refresh token
func GenerateRefreshToken(userID string, rememberMe bool, tx *gorm.DB) (string, error) {
	lifespan, err := GetRefreshTime(rememberMe)
	if err != nil {
		return "", err
	}
	jti, err := gonanoid.New()
	if err != nil {
		return "", err
	}
	exp := time.Now().Add(time.Hour * time.Duration(lifespan))
	userSession := &data.UserSession{
		UserID:     userID,
		JTI:        jti,
		ExpireDate: &exp,
		RememberMe: rememberMe,
	}
	err = userSession.SaveSession(tx)
	if err != nil {
		return "", err
	}
	claims := jwt.MapClaims{}
	claims["jti"] = jti
	claims["sub"] = userID
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(lifespan)).Unix()
	claims["rememberMe"] = rememberMe
	claims["type"] = "refresh"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
}

// Exctract access token from cookie or Authorization header
func ExtractToken(c *gin.Context) string {
	if access_token, err := c.Request.Cookie("token"); err == nil {
		return access_token.Value
	}
	if bearerToken := c.Request.Header.Get("Authorization"); bearerToken != "" && len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	if queryToken := c.Query("token"); queryToken != "" {
		return queryToken
	}
	return ""
}

// Validate token
func ValidateToken(tokenIn string) (token *jwt.Token, err error) {
	token, err = jwt.Parse(tokenIn, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// Extract payload from token
func ExtractTokenClaims(token *jwt.Token) (jwt.MapClaims, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}

// Get access token time from .env
func GetAccessTime() (int, error) {
	return strconv.Atoi(os.Getenv("TOKEN_TIME_ACCESS"))
}

// Get refresh token time from .env
func GetRefreshTime(rememberMe bool) (int, error) {
	if rememberMe {
		return strconv.Atoi(os.Getenv("TOKEN_REMEMBER_REFRESH"))
	}
	return strconv.Atoi(os.Getenv("TOKEN_TIME_REFRESH"))
}
