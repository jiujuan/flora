package lilac

import (
	"html/template"

	"github.com/julienschmidt/httprouter"
)

type App struct {
	*Router
	router   *httprouter.Router
	Template *template.Template
}

func New() *App {
	app := new(App)
	app.router = httprouter.New()
	return app
}
