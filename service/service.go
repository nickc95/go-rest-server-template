package service

import (
	"github.com/julienschmidt/httprouter"
	"net/http"

	"go-rest-server-template/models"
)

var serviceRoutePrefix = "/service/"

// New defines identity service package, including service specific middleware and handlers
func New(router *httprouter.Router, wrapper func(http.HandlerFunc) httprouter.Handle, commonMiddlewareWrapper func(func(*models.ResponseContextWriter, *http.Request, *models.RequestData)) http.HandlerFunc) {
	router.POST(serviceRoutePrefix, wrapper(commonMiddlewareWrapper(serviceMiddleware(servicePostHandler))))
	router.GET(serviceRoutePrefix, wrapper(commonMiddlewareWrapper(serviceMiddleware(serviceGetHandler))))
}

func serviceMiddleware(fn func(*models.ResponseContextWriter, *http.Request, *models.RequestData)) func(*models.ResponseContextWriter, *http.Request, *models.RequestData) {
	return func(w *models.ResponseContextWriter, r *http.Request, requestData *models.RequestData) {
		fn(w, r, requestData)
	}
}

func servicePostHandler(w *models.ResponseContextWriter, r *http.Request, requestData *models.RequestData) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func serviceGetHandler(w *models.ResponseContextWriter, r *http.Request, requestData *models.RequestData) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
