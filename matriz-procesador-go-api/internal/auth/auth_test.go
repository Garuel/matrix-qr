package auth

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testSecret = "test-secret-key"

func TestLogin(t *testing.T) {
	t.Parallel()

	t.Run("credenciales válidas retorna token JWT válido", func(t *testing.T) {
		t.Parallel()
		service := NewAuthService(testSecret)

		tokenString, err := service.Login("admin", "admin123")

		require.NoError(t, err)
		assert.NotEmpty(t, tokenString)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(testSecret), nil
		})
		require.NoError(t, err)
		assert.True(t, token.Valid)
	})

	t.Run("token contiene claims correctos", func(t *testing.T) {
		t.Parallel()
		service := NewAuthService(testSecret)

		tokenString, err := service.Login("admin", "admin123")
		require.NoError(t, err)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(testSecret), nil
		})
		require.NoError(t, err)

		claims, ok := token.Claims.(jwt.MapClaims)
		require.True(t, ok, "los claims deben ser de tipo MapClaims")

		assert.Equal(t, "admin", claims["user"], "claim 'user' debe ser 'admin'")
		assert.Equal(t, "admin", claims["role"], "claim 'role' debe ser 'admin'")
	})

	t.Run("token tiene expiración de 2 horas en el futuro", func(t *testing.T) {
		t.Parallel()
		service := NewAuthService(testSecret)

		before := time.Now()
		tokenString, err := service.Login("admin", "admin123")
		require.NoError(t, err)
		after := time.Now()

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(testSecret), nil
		})
		require.NoError(t, err)

		claims, ok := token.Claims.(jwt.MapClaims)
		require.True(t, ok)

		exp, ok := claims["exp"].(float64)
		require.True(t, ok, "claim 'exp' debe existir y ser numérico")

		expTime := time.Unix(int64(exp), 0)
		// Tolerancia de 2 segundos por truncamiento de Unix() y overhead de bcrypt
		expectedMin := before.Add(2 * time.Hour).Add(-2 * time.Second)
		expectedMax := after.Add(2 * time.Hour).Add(2 * time.Second)

		assert.True(t, !expTime.Before(expectedMin) && !expTime.After(expectedMax),
			"la expiración debe estar entre %v y %v, pero fue %v", expectedMin, expectedMax, expTime)
	})

	t.Run("token usa método de firma HMAC", func(t *testing.T) {
		t.Parallel()
		service := NewAuthService(testSecret)

		tokenString, err := service.Login("admin", "admin123")
		require.NoError(t, err)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			assert.True(t, ok, "el método de firma debe ser HMAC")
			return []byte(testSecret), nil
		})
		require.NoError(t, err)
		assert.True(t, token.Valid)
	})
}

func TestLogin_CredencialesInvalidas(t *testing.T) {
	t.Parallel()

	tests := []struct {
		nombre   string
		usuario  string
		password string
	}{
		{
			nombre:   "usuario incorrecto",
			usuario:  "usuarioInvalido",
			password: "admin123",
		},
		{
			nombre:   "password incorrecta",
			usuario:  "admin",
			password: "passwordIncorrecta",
		},
		{
			nombre:   "ambos incorrectos",
			usuario:  "hacker",
			password: "12345",
		},
		{
			nombre:   "usuario vacío",
			usuario:  "",
			password: "admin123",
		},
		{
			nombre:   "password vacía",
			usuario:  "admin",
			password: "",
		},
		{
			nombre:   "ambos vacíos",
			usuario:  "",
			password: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.nombre, func(t *testing.T) {
			t.Parallel()
			service := NewAuthService(testSecret)

			token, err := service.Login(tc.usuario, tc.password)

			assert.Error(t, err)
			assert.Equal(t, "credenciales inválidas", err.Error())
			assert.Empty(t, token, "no se debe retornar token con credenciales inválidas")
		})
	}
}

func TestLogin_DiferentesSecrets(t *testing.T) {
	t.Parallel()

	t.Run("token firmado con un secret no se valida con otro", func(t *testing.T) {
		t.Parallel()
		service := NewAuthService("secret-original")

		tokenString, err := service.Login("admin", "admin123")
		require.NoError(t, err)

		// Intentar parsear con un secret diferente debe fallar
		_, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret-diferente"), nil
		})
		assert.Error(t, err, "el token no debe validarse con un secret diferente")
	})

	t.Run("diferentes secrets generan diferentes tokens", func(t *testing.T) {
		t.Parallel()
		service1 := NewAuthService("secret-1")
		service2 := NewAuthService("secret-2")

		token1, err := service1.Login("admin", "admin123")
		require.NoError(t, err)

		token2, err := service2.Login("admin", "admin123")
		require.NoError(t, err)

		assert.NotEqual(t, token1, token2,
			"tokens generados con secrets diferentes deben ser distintos")
	})
}
