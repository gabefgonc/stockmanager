package products

import "github.com/gabefgonc/stockmanager/internal/infra/models"

type Product struct {
	models.BaseModel
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Amount      uint    `json:"amount"`
}
