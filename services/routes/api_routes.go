package routes

import (
	"github.com/faber-numeris/cirrus/spec/openapi"
	"github.com/sirupsen/logrus"
	"io/fs"
	"net/http"
)

func (router RouterImpl) LambdaFunction(w http.ResponseWriter, r *http.Request, functionName string) {
	w.WriteHeader(http.StatusOK)
	var staticFS = fs.FS(openapi.Files)
	htmlContent, err := fs.Sub(staticFS, "files")
	if err != nil {
		logrus.Error("Error reading openapi spec server files content", err.Error())
		return
	}
	specServer := http.FileServer(http.FS(htmlContent))
	specServer.ServeHTTP(w, r)
}
