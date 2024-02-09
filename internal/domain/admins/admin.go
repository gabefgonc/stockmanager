package admins

import "github.com/gabefgonc/stockmanager/internal/infra/models"

type Admin struct {
	models.BaseModel
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"-"`
}
