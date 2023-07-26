package app

import (
	"net/http"

	"github.com/gocopper/copper/chttp"
	"github.com/gocopper/copper/clogger"
)

type NewRouterParams struct {
	RW     *chttp.ReaderWriter
	Logger clogger.Logger
}

func NewRouter(p NewRouterParams) *Router {
	return &Router{
		rw:     p.RW,
		logger: p.Logger,
	}
}

type Router struct {
	rw     *chttp.ReaderWriter
	logger clogger.Logger
}

func (ro *Router) Routes() []chttp.Route {
	return []chttp.Route{
		{
			Path:    "/",
			Methods: []string{http.MethodGet},
			Handler: ro.HandleIndexPage,
		},
		{
			Path:    "/menu",
			Methods: []string{http.MethodGet},
			Handler: ro.HandleMenuPage,
		},
	}
}

func (ro *Router) HandleIndexPage(w http.ResponseWriter, r *http.Request) {

	ro.rw.WriteHTML(w, r, chttp.WriteHTMLParams{
		PageTemplate: "index.html",
	})
}

func (ro *Router) HandleMenuPage(w http.ResponseWriter, r *http.Request) {

	ro.rw.WriteHTML(w, r, chttp.WriteHTMLParams{
		PageTemplate: "menu.html",
	})
}
