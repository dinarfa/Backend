package middleware

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

// CreateTokenJWT generates a JWT token for a user
func CreateTokenJWT(userID int) (string, error) {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the secret key from environment variables
	secretKey := []byte(os.Getenv("JWT_KEY")) // Convert to []byte

	// Create claims
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Minute * 5).Unix(),
	}

	// Create token
	tokenJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := tokenJWT.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateTokenJWT validates the JWT token and returns its claims
func ValidateTokenJWT(tokenString string, secretKey []byte) (jwt.MapClaims, error) {
	// Parse the token
	tokenJWT, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		// Ensure token uses HMAC signing method
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return secretKey, nil
	})
	if err != nil {
		return nil, errors.New("invalid token")
	}

	// Validate the token and return claims
	if claims, ok := tokenJWT.Claims.(jwt.MapClaims); ok && tokenJWT.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// JWTMiddleware is the middleware function for validating JWT tokens
func JWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Load environment variables
		err := godotenv.Load()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error loading .env file"})
			ctx.Abort()
			return
		}

		// Get the secret key from environment variables
		secretKey := []byte(os.Getenv("JWT_KEY")) // Convert to []byte

		// Get the Authorization header
		headerToken := ctx.GetHeader("Authorization")
		if headerToken == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			return
		}

		// Check if token has Bearer prefix
		if !strings.HasPrefix(headerToken, "Bearer ") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})

			return
		}

		// Extract the token string
		tokenString := strings.TrimPrefix(headerToken, "Bearer ")

		// Validate the token
		claims, err := ValidateTokenJWT(tokenString, secretKey)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		// Set claims to context for further use
		ctx.Set("claims", claims)
		ctx.Next()
	}
}