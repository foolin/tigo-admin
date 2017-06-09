package routes

import (
	"github.com/foolin/tigo"
	"strings"
	"fmt"
	"net/http"
)

func AdminRoute(ctx *tigo.Context) error {
	//check login
	username := ctx.GetCookieValue("username")
	if username == "" {
		ctx.Redirect("/login", http.StatusOK)
		return nil
	}

	//admin page
	id := ctx.Param("id")
	id = strings.TrimRight(strings.ToLower(id), ".html")

	switch id {
	case "":
		id = "index"
		fallthrough
	case "index", "charts", "forms", "icons", "panels", "tables", "widgets":
		return ctx.Render(fmt.Sprintf("admin/%v", id), tigo.M{
			"Id": id,
			"Username": username,
		})
	}

	return ctx.NotFound()
}