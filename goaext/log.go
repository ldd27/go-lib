package goaext

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/goadesign/goa"
)

// LogRequest creates a request logger middleware.
// This middleware is aware of the RequestID middleware and if registered after it leverages the
// request ID for logging.
// If verbose is true then the middlware logs the request and response bodies.
func LogRequest(verbose bool) goa.Middleware {
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {

			if strings.Contains(req.URL.Path, "/health") {
				return h(ctx, rw, req)
			}

			reqID := ctx.Value(ReqIDKey)
			if reqID == nil {
				reqID = shortID()
			}
			ctx = goa.WithLogContext(ctx, "req_id", reqID)
			startedAt := time.Now()
			r := goa.ContextRequest(ctx)
			goa.LogInfo(ctx, "started", "request", fmt.Sprintf("%s %s", r.Method, r.URL.String()), "from", from(req),
				"ctrl", goa.ContextController(ctx), "action", goa.ContextAction(ctx))
			if verbose {
				if len(r.Header) > 0 {
					logCtx := make([]interface{}, 2*len(r.Header))
					i := 0
					for k, v := range r.Header {
						logCtx[i] = k
						logCtx[i+1] = interface{}(strings.Join(v, ", "))
						i = i + 2
					}
					goa.LogInfo(ctx, "headers", logCtx...)
				}
				if len(r.Params) > 0 {
					logCtx := make([]interface{}, 2*len(r.Params))
					i := 0
					for k, v := range r.Params {
						logCtx[i] = k
						logCtx[i+1] = interface{}(strings.Join(v, ", "))
						i = i + 2
					}
					goa.LogInfo(ctx, "params", logCtx...)
				}
				if r.ContentLength > 0 {
					if mp, ok := r.Payload.(map[string]interface{}); ok {
						logCtx := make([]interface{}, 2*len(mp))
						i := 0
						for k, v := range mp {
							logCtx[i] = k
							logCtx[i+1] = interface{}(v)
							i = i + 2
						}
						goa.LogInfo(ctx, "payload", logCtx...)
					} else {
						// Not the most efficient but this is used for debugging
						js, err := json.Marshal(r.Payload)
						if err != nil {
							js = []byte("<invalid JSON>")
						}
						goa.LogInfo(ctx, "payload", "raw", string(js))
					}
				}
			}
			strReqID, _ := reqID.(string)
			rw.Header().Set("x-req-id", strReqID)
			err := h(ctx, rw, req)
			if err != nil {
				switch vlu := err.(type) {
				case *goa.ErrorResponse:
					goa.LogInfo(ctx, "completed", "status", vlu.ResponseStatus(), "time", time.Since(startedAt).String(),
						"ctrl", goa.ContextController(ctx), "action", goa.ContextAction(ctx), "error", vlu.Detail)
				default:
					goa.LogInfo(ctx, "completed", "time", time.Since(startedAt).String(),
						"ctrl", goa.ContextController(ctx), "action", goa.ContextAction(ctx), "error", vlu.Error())
				}
			} else {
				resp := goa.ContextResponse(ctx)
				if resp != nil {
					goa.LogInfo(ctx, "completed", "status", resp.Status, "bytes", resp.Length, "time", time.Since(startedAt).String(),
						"ctrl", goa.ContextController(ctx), "action", goa.ContextAction(ctx))
				} else {
					goa.LogInfo(ctx, "completed", time.Since(startedAt).String(),
						"ctrl", goa.ContextController(ctx), "action", goa.ContextAction(ctx), "resp", "null")
				}
			}
			return err
		}
	}
}

// shortID produces a "unique" 6 bytes long string.
// Do not use as a reliable way to get unique IDs, instead use for things like logging.
func shortID() string {
	b := make([]byte, 6)
	io.ReadFull(rand.Reader, b)
	return base64.StdEncoding.EncodeToString(b)
}

// from makes a best effort to compute the request client IP.
func from(req *http.Request) string {
	if f := req.Header.Get("X-Forwarded-For"); f != "" {
		return f
	}
	f := req.RemoteAddr
	ip, _, err := net.SplitHostPort(f)
	if err != nil {
		return f
	}
	return ip
}
