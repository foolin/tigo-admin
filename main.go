package main

import (
	"github.com/foolin/tigo"
	"html/template"
	"log"
	"github.com/foolin/tigo-admin/routes"
)

func main() {
	router := tigo.Default()

	router.SetRender(tigo.NewViewRender(tigo.ViewRenderConfig{
		Root: "views",
		Extension: ".html",
		Master: "layout/master",
		Partials: []string{},
		Funcs: template.FuncMap{
		},
		DisableCache: true, //debug mode, use true
		DisableFilePartial: false,
	}))

	router.Get("/static/*", tigo.Static("static/", 1))

	//404
	router.OnNotFound = func(ctx *tigo.Context) error {
		return ctx.RenderFile("404", tigo.M{})
	}
	//500
	router.OnError = func(ctx *tigo.Context, err error) {
		err = ctx.RenderFile("500", tigo.M{
			"Msg": err.Error(),
		})
		if err != nil {
			ctx.HTML(err.Error())
		}
	}
	//routes
	router.Any("/", routes.IndexRoute)
	router.Any("/login", routes.LoginRoute)
	router.Any("/logout", routes.LogoutRoute)
	router.Any("/admin/<id>", routes.AdminRoute)

	log.Printf("run on: 8080")
	log.Fatal(router.Run(":8080"))
}

