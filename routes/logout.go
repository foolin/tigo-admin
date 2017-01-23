package routes

import "github.com/foolin/tigo"

func LogoutRoute(ctx *tigo.Context) error {
	//delete cookie
	ctx.DelCookie("username")

	//jump to login
	ctx.Redirect("/login")

	return nil
}
