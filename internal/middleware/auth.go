package middleware

import (
	"net/http"

	response "rest_base/internal/category/web/response"
	"rest_base/pkg/logger"
	"rest_base/pkg/utils"

	"github.com/julienschmidt/httprouter"
)

var AuthMiddleware = func(f httprouter.Handle, _ logger.Logger) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		if "SECRET" == r.Header.Get("X-API-Key") {
			f(w, r, params)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)

			webResponse := response.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "UNAUTHORIZED",
			}

			utils.WriteToResponseBody(w, webResponse)
		}
	}
}
