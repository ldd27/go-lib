package goaext

import (
	"context"
	"net/http"

	"github.com/goadesign/goa"
)

func NewCORSMiddleware() goa.Middleware {
	return CORSMiddleware
}

func CORSMiddleware(h goa.Handler) goa.Handler {
	return func(ctx context.Context, writer http.ResponseWriter, request *http.Request) error {
		origin := request.Header.Get("Origin")
		if origin != "" {
			writer.Header().Set("Access-Control-Allow-Origin", origin)
			writer.Header().Set("Access-Control-Allow-Methods", "HEAD, OPTIONS, GET, POST, PUT, DELETE")
			writer.Header().Set("Access-Control-Max-Age", "86400")
			writer.Header().Set("Access-Control-Allow-Credentials", "true")
			writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		}
		return h(ctx, writer, request)
	}
}
