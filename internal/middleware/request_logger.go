package middleware

import (
	"golang_restfull_api/pkg/logger"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

var RequestLoggerMiddleware = func(f httprouter.Handle, logger logger.Logger) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		start := time.Now()
		duration := time.Since(start).String()
		logger.Infof("Method: %s, URI: %s, Time: %s", r.Method, r.URL, duration)

		f(w, r, params)
	}
}
