package server

import (
	"net/http"
	"strings"

	"go-rest-server-template/models"
)

// normalizeParamCasing normalizes request query params, storing them in requestData.RequestQuery map
func normalizeParamCasing(w *models.ResponseContextWriter, r *http.Request, requestData *models.RequestData) {
	query := r.URL.Query()

	for key, _ := range query {
		requestData.RequestQuery[strings.ToLower(key)] = query.Get(key)
	}
}
