package routes

import "gopkg.in/foolin/tigo.v1"

func LogoutRoute(ctx *tigo.Context) error {
	//delete cookie
	ctx.DelCookie("username")

	//jump to login
	ctx.Redirect("/login")

	return nil
}
