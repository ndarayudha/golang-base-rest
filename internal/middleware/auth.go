package middleware

import (
	"golang_restfull_api/internal/category/web"
	"golang_restfull_api/pkg/logger"
	"golang_restfull_api/pkg/utils"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var AuthMiddleware = func(f httprouter.Handle, _ logger.Logger) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		if "SECRET" == r.Header.Get("X-API-Key") {
			f(w, r, params)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)

			webResponse := web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "UNAUTHORIZED",
			}

			utils.WriteToResponseBody(w, webResponse)
		}
	}
}
