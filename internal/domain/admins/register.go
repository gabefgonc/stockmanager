package admins

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterAdmin struct {
	DB *gorm.DB
}

func (r *RegisterAdmin) Execute(name string, email string, phoneNumber string, password string) (*Admin, error) {
	if len(name) < 5 {
		return nil, ErrNameTooShort
	}

	if len(name) > 50 {
		return nil, ErrNameTooLong
	}

	const emailRegex = "/^[a-z0-9.]+@[a-z0-9]+\\.[a-z]+\\.([a-z]+)?$/i"
	emailValid, _ := regexp.MatchString(emailRegex, email)

	if !emailValid {
		return nil, ErrEmailInvalid
	}

	if len(phoneNumber) < 11 {
		return nil, ErrPhoneNumberInvalid
	}

	if len(password) < 6 {
		return nil, ErrPasswordTooShort
	}

	// Password limit so bcrypt.ErrPasswordTooLong is not thrown
	if len(password) > 70 {
		return nil, ErrPasswordTooLong
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	admin := &Admin{
		Name:        name,
		Email:       email,
		PhoneNumber: phoneNumber,
		Password:    string(hashPassword),
	}

	result := r.DB.Create(admin)
	if result.Error != nil {
		return nil, ErrInternal
	}

	return admin, nil
}

func NewRegisterAdmin(db *gorm.DB) *RegisterAdmin {
	return &RegisterAdmin{
		DB: db,
	}
}
