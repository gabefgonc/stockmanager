package admins

import "go.uber.org/fx"

var Module = fx.Provide(
	NewAuthenticateAdmin,
	NewRegisterAdmin,
)
