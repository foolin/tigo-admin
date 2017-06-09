package routes

import (
	"github.com/foolin/tigo"
	"net/http"
)

func IndexRoute(ctx *tigo.Context) error {
	ctx.Redirect("/login", http.StatusOK)
	return nil
}