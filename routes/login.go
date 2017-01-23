package routes

import (
	"github.com/foolin/tigo"
	"time"
)

func LoginRoute(ctx *tigo.Context) error {
	//post
	if ctx.IsPost(){
		username := ctx.FormString("username")
		if username != ""{
			ctx.SetCookieValue("username", username, time.Now().Add(time.Hour))
			ctx.Redirect("/admin/index")
			return nil
		}
	}

	//render view
	return ctx.RenderFile("login", tigo.M{})
}
