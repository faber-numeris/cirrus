package routes

import (
	api "github.com/faber-numeris/cirrus/api/rest"
	"github.com/faber-numeris/cirrus/spec/openapi"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"io/fs"
	"log"
	"net/http"
)

type Router interface {
}

type RouterImpl struct {
}

func NewRouter() *RouterImpl {
	return &RouterImpl{}
}

func (router RouterImpl) Run() {

	chiMux := chi.NewMux()
	h := api.HandlerFromMux(router, chiMux)
	s := &http.Server{
		Handler: h,
		Addr:    "0.0.0.0:8080",
	}

	logrus.Info("Starting cirrus server on port 8080")
	logrus.Info("Available routes:")
	err := chi.Walk(chiMux, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		logrus.Infof("%s\t%s\n", method, route) // Walk and print out all routes
		return nil
	})
	if err != nil {
		return
	}

	// OpenAPI Spec
	htmlContent, err := fs.Sub(openapi.Files, ".")
	if err != nil {
		logrus.Error("Error serving OpenAPI Spec", err)
	}
	openapiServer := http.FileServer(http.FS(htmlContent))
	chiMux.Handle("/openapi/*", http.StripPrefix("/openapi/", openapiServer))
	logrus.Infof("GET\t/openapi/")

	// Websocket
	chiMux.Handle("/ws", http.HandlerFunc(HandleWebSocket))

	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())

}
