package controller

import (
	"fmt"
	"net/http"

	"github.com/wcharczuk/go-web"
	"github.com/wcharczuk/instabot/server/core"
)

// Index is the root controller.
type Index struct{}

func (i Index) indexAction(r *web.RequestContext) web.ControllerResult {
	return r.View().View("index", nil)
}

// Register registers the controller
func (i Index) Register(app *web.App) {
	app.GET("/", i.indexAction)
	app.GET("/index.html", i.indexAction)

	if core.ConfigIsProduction() {
		app.Static("/static/*filepath", http.Dir("_client/dist"))
		app.StaticRewrite("/static/*filepath", `^(.*)\.([0-9]+)\.(css|js)$`, func(path string, parts ...string) string {
			if len(parts) < 4 {
				return path
			}
			return fmt.Sprintf("%s.%s", parts[1], parts[3])
		})
		app.StaticHeader("/static/*filepath", "access-control-allow-origin", "*")
		app.StaticHeader("/static/*filepath", "cache-control", "public,max-age=315360000")
	} else {
		app.Static("/bower/*filepath", http.Dir("_client/bower"))
		app.Static("/static/*filepath", http.Dir("_client/src"))
	}
}
