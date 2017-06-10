package main

import (
	"html/template"
	"log"
	"github.com/foolin/tigo-admin/routes"
	"gopkg.in/foolin/tigo.v2"
)

func main() {
	router := tigo.Default()

	router.Render = tigo.NewViewRender(tigo.ViewRenderConfig{
		Root: "views",
		Extension: ".html",
		Master: "layout/master",
		Partials: []string{},
		Funcs: template.FuncMap{
		},
		DisableCache: true, //debug mode, use true
		DisableFilePartial: false,
	})

	router.File("/favicon.ico", "static/favicon.ico")
	router.Static("/static/*", "./static")

	//404
	router.NotFound(func(ctx *tigo.Context) error {
		return ctx.RenderFile("404", tigo.M{})
	})
	//routes
	router.Any("/", routes.IndexRoute)
	router.Any("/login", routes.LoginRoute)
	router.Any("/logout", routes.LogoutRoute)
	router.Any("/admin/<id>", routes.AdminRoute)

	log.Printf("run on: 8080")
	log.Fatal(router.Run(":8080"))
}

