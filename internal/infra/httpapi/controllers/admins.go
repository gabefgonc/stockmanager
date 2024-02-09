package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gabefgonc/stockmanager/internal/domain/admins"
	"github.com/go-chi/render"
)

type AdminsController struct {
	registerAdmin     *admins.RegisterAdmin
	authenticateAdmin *admins.AuthenticateAdmin
}

type RegisterAdminDTO struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

func (a *AdminsController) HandleRegisterAdmin(w http.ResponseWriter, r *http.Request) {
	var dto RegisterAdminDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, render.M{})
		return
	}

	admin, err := a.registerAdmin.Execute(dto.Name, dto.Email, dto.PhoneNumber, dto.Password)
	if err != nil {
		if !errors.Is(err, admins.ErrInternal) {
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, render.M{
				"error": err.Error(),
			})
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, render.M{
				"error": err.Error(),
			})
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, render.M{
		"admin": admin,
	})
}

type AuthenticateAdminDTO struct {
	Email    string
	Password string
}

func (a *AdminsController) AuthenticateAdmin(w http.ResponseWriter, r *http.Request) {
	var dto AuthenticateAdminDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, render.M{})
		return
	}

	token, err := a.authenticateAdmin.Execute(dto.Email, dto.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		render.JSON(w, r, render.M{"error": "wrong credentials"})
		return
	}

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, render.M{
		"token": token,
	})
}

func NewAdminsController(registerAdmin *admins.RegisterAdmin, authenticateAdmin *admins.AuthenticateAdmin) *AdminsController {
	return &AdminsController{
		registerAdmin:     registerAdmin,
		authenticateAdmin: authenticateAdmin,
	}
}
