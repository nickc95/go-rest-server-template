package server

import (
	"github.com/julienschmidt/httprouter"
	"net/http"

	"go-rest-server-template/config"
	"go-rest-server-template/errors"
	"go-rest-server-template/models"
	"go-rest-server-template/service"
)

var configObj *config.ConfigObject

func init() {
	configObj = config.GetInstance()
	loggerInit()
}

// NewServerRouter creates new server router, composing service packages together
func NewServerRouter() *httprouter.Router {
	router := httprouter.New()

	router.HandleMethodNotAllowed = false
	router.NotFound = http.HandlerFunc(notFound)

	service.New(router, wrapHandler, defineCommonMiddleware)

	return router
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	responseJSON := errors.GetErrorJSON(errors.InternalServerError)
	w.Write(responseJSON)
}

// wrapHandler wraps http.HandlerFunc and returns httprouter.Handle
func wrapHandler(h http.HandlerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := r.Context()
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	}
}

func defineCommonMiddleware(fn func(*models.ResponseContextWriter, *http.Request, *models.RequestData)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Initialize contextwriter and requestdata structs to record and pass data between middleware/service/logger
		responseContextWriter := models.ResponseContextWriter{ResponseWriter: w}
		requestData := models.NewRequestData()

		normalizeParamCasing(&responseContextWriter, r, requestData)

		fn(&responseContextWriter, r, requestData)
		if responseContextWriter.IsError == true {
			log(responseContextWriter, r, requestData)
		}
	}
}
