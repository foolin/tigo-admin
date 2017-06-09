package routes

import (
	"github.com/foolin/tigo"
)

func IndexRoute(ctx *tigo.Context) error {
	ctx.Redirect("/login")
	return nil
}