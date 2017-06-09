package routes

import (
	"gopkg.in/foolin/tigo.v2"
	"time"
)

func LoginRoute(ctx *tigo.Context) error {
	//post
	if ctx.IsPost(){
		username := ctx.Request.FormValue("username")
		if username != ""{
			ctx.SetCookieValue("username", username, time.Now().Add(time.Hour))
			ctx.Redirect("/admin/index")
			return nil
		}
	}

	//render view
	return ctx.RenderFile("login", tigo.M{})
}
