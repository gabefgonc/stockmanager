package stores

import (
	"github.com/gabefgonc/stockmanager/internal/domain/admins"
	"github.com/gabefgonc/stockmanager/internal/domain/stocks"
	"github.com/gabefgonc/stockmanager/internal/infra/models"
)

type Store struct {
	models.BaseModel
	Name    string         `json:"name"`
	Address string         `json:"address"`
	ZipCode string         `json:"zipCode"`
	Stock   stocks.Stock   `json:"stock"`
	Admins  []admins.Admin `json:"admins"`
}
