package auth

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Login(user, pass string) (string, error)
}

type AuthService struct {
	secretKey string
}

func NewAuthService(secret string) IAuthService {
	return &AuthService{secretKey: secret}
}

func (s *AuthService) Login(user, pass string) (string, error) {
	usuario := "admin"
	
	    log.Println("Creando hash de password...")
	passhash, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)

	if user != usuario {
		return "", errors.New("credenciales inválidas")
	}

	if err := bcrypt.CompareHashAndPassword(passhash, []byte(pass)); err != nil {
		return "", errors.New("credenciales inválidas")
	}

	log.Printf("Credenciales válidas")

	log.Printf("Generando token JWT...")

	claims := jwt.MapClaims{
		"user": user,
		"role": "admin",
		"exp":  time.Now().Add(time.Hour * 2).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	log.Printf("Token JWT generado exitosamente")

	return token.SignedString([]byte(s.secretKey))
}