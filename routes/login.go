package routes

import (
	"github.com/foolin/tigo"
	"time"
	"net/http"
)

func LoginRoute(ctx *tigo.Context) error {
	//post
	if ctx.IsPost(){
		username := ctx.Request.FormValue("username")
		if username != ""{
			ctx.SetCookieValue("username", username, time.Now().Add(time.Hour))
			ctx.Redirect("/admin/index", http.StatusOK)
			return nil
		}
	}

	//render view
	return ctx.RenderFile("login", tigo.M{})
}
