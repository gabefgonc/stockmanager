package httpapi

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gabefgonc/stockmanager/internal/infra/httpapi/routers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func NewHTTPServer(lc fx.Lifecycle, adminsRouter *routers.AdminsRouter) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {

			err := godotenv.Load()
			if err != nil {
				panic("dotenv fails to load")
			}

			r := chi.NewRouter()
			r.Use(
				render.SetContentType(render.ContentTypeJSON),
				middleware.Logger,
			)

			r.Mount("/admins", adminsRouter.Routes())

			port := os.Getenv("PORT")
			go http.ListenAndServe(fmt.Sprintf(":%s", port), r)
			return nil
		},
	})
}
