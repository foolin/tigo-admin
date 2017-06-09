package routes

import (
	"gopkg.in/foolin/tigo.v2"
)

func IndexRoute(ctx *tigo.Context) error {
	ctx.Redirect("/login")
	return nil
}