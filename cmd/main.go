package main

import (
	"github.com/gabefgonc/stockmanager/internal/domain/admins"
	"github.com/gabefgonc/stockmanager/internal/infra/database"
	"github.com/gabefgonc/stockmanager/internal/infra/httpapi"
	"github.com/gabefgonc/stockmanager/internal/infra/httpapi/controllers"
	"github.com/gabefgonc/stockmanager/internal/infra/httpapi/routers"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		controllers.Module,
		routers.Module,
		database.Module,
		admins.Module,
		fx.Invoke(
			httpapi.NewHTTPServer,
		),
	).Run()
}
