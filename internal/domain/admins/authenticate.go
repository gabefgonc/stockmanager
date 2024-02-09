package admins

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthenticateAdmin struct {
	DB *gorm.DB
}

func (a *AuthenticateAdmin) Execute(email string, password string) (string, error) {

	var admin Admin
	a.DB.Model(&Admin{}).Where("email = ?", email).Find(&admin)

	if admin.Email == "" {
		return "", ErrWrongCredentials
	}

	err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))

	if err != nil {
		return "", ErrWrongCredentials
	}

	key := []byte(os.Getenv("APP_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 3)),
		Subject:   email,
	})

	signedString, err := token.SignedString(key)
	return signedString, err
}

func NewAuthenticateAdmin(db *gorm.DB) *AuthenticateAdmin {
	return &AuthenticateAdmin{
		DB: db,
	}
}
