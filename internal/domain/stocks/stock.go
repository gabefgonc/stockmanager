package stocks

import (
	"github.com/gabefgonc/stockmanager/internal/domain/products"
	"github.com/gabefgonc/stockmanager/internal/infra/models"
)

type Stock struct {
	models.BaseModel
	Name     string             `json:"name"`
	Products []products.Product `json:"products"`
}
