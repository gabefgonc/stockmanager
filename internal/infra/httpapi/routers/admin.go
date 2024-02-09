package routers

import (
	"github.com/gabefgonc/stockmanager/internal/infra/httpapi/controllers"
	"github.com/go-chi/chi/v5"
)

type AdminsRouter struct {
	adminsController *controllers.AdminsController
}

func (a *AdminsRouter) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", a.adminsController.HandleRegisterAdmin)

	return r
}

func NewAdminsRouter(adminsController *controllers.AdminsController) *AdminsRouter {
	return &AdminsRouter{}
}
