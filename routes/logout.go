package routes

import (
	"github.com/foolin/tigo"
	"net/http"
)

func LogoutRoute(ctx *tigo.Context) error {
	//delete cookie
	ctx.DelCookie("username")

	//jump to login
	ctx.Redirect("/login", http.StatusOK)

	return nil
}
