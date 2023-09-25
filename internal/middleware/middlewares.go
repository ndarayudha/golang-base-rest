package middleware

import (
	"rest_base/pkg/logger"

	"github.com/julienschmidt/httprouter"
)

// middleware is a definition of  what a middleware is,
// take in one handlerfunc and wrap it within another handlerfunc
type Middlewares func(httprouter.Handle, logger.Logger) httprouter.Handle

// buildChain builds the middlware chain recursively, functions are first class
func BuildChain(f httprouter.Handle, log logger.Logger, m ...Middlewares) httprouter.Handle {
	// if our chain is done, use the original handlerfunc
	if len(m) == 0 {
		return f
	}
	// otherwise nest the handlerfuncs
	return m[0](BuildChain(f, log, m[1:cap(m)]...), log)
}
