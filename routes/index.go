package routes

import "gopkg.in/foolin/tigo.v1"

func IndexRoute(ctx *tigo.Context) error {
	ctx.Redirect("/login")
	return nil
}