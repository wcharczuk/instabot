package server

import (
	"log"

	"github.com/wcharczuk/go-web"
	"github.com/wcharczuk/instabot/server/controller"
	"github.com/wcharczuk/instabot/server/core"
)

func headerPath() string {
	if core.ConfigIsProduction() {
		return "server/_views/partials/header_prod.html"
	}
	return "server/_views/partials/header.html"
}

func footerPath() string {
	return "server/_views/partials/footer.html"
}

func viewPaths() []string {
	return []string{
		headerPath(),
		footerPath(),
		"server/_views/error/not_found.html",
		"server/_views/error/error.html",
		"server/_views/error/bad_request.html",
		"server/_views/error/not_authorized.html",
		"server/_views/index.html",
	}
}

// Init inits the app.
func Init() *web.App {
	core.DBInit()

	app := web.New()
	app.SetName("instabot")
	app.SetPort(core.ConfigPort())

	viewCacheErr := app.InitViewCache(viewPaths()...)
	if viewCacheErr != nil {
		log.Fatal(viewCacheErr)
	}

	if !core.ConfigIsProduction() {
		app.SetLogger(web.NewStandardOutputLogger())
	}

	app.Register(new(controller.Index))

	return app
}
