package auth

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestAuthService_Login_Success(t *testing.T) {
	secret := "my-secret-key"
	service := NewAuthService(secret)

	tokenString, err := service.Login("admin", "admin123")

	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)

	// Validar el token generado
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	assert.NoError(t, err)
	assert.True(t, token.Valid)

	claims, ok := token.Claims.(jwt.MapClaims)
	assert.True(t, ok)
	assert.Equal(t, "admin", claims["user"])
	assert.Equal(t, "admin", claims["role"])
	
	// Validar expiración
	exp, ok := claims["exp"].(float64)
	assert.True(t, ok)
	assert.True(t, exp > float64(time.Now().Unix()))
}

func TestAuthService_Login_InvalidUser(t *testing.T) {
	service := NewAuthService("secret")

	token, err := service.Login("wronguser", "admin123")

	assert.Error(t, err)
	assert.Equal(t, "credenciales inválidas", err.Error())
	assert.Empty(t, token)
}

func TestAuthService_Login_InvalidPassword(t *testing.T) {
	service := NewAuthService("secret")

	token, err := service.Login("admin", "wrongpass")

	assert.Error(t, err)
	assert.Equal(t, "credenciales inválidas", err.Error())
	assert.Empty(t, token)
}
