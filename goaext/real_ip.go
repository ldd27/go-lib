package goaext

import (
	"context"
	"net"
	"net/http"
	"strings"

	"github.com/goadesign/goa"
)

func NewRealIPMiddleware() goa.Middleware {
	return RealIPMiddleware
}

func RealIPMiddleware(h goa.Handler) goa.Handler {
	return func(ctx context.Context, writer http.ResponseWriter, request *http.Request) error {
		realIP := GetRealIP(request)
		ctx = context.WithValue(ctx, RealIPKey, realIP)
		return h(ctx, writer, request)
	}
}

// 获取真实IP
func GetRealIP(req *http.Request) string {
	clientIP := req.Header.Get("Ali-Cdn-Real-Ip") // real ip from aliyun cdn
	if clientIP != "" {
		return clientIP
	}
	clientIP = req.Header.Get("Remoteip") // real ip from aliyun slb http proxy
	if clientIP != "" {
		return clientIP
	}
	clientIP = req.Header.Get("X-Forwarded-For")
	if clientIP != "" {
		return strings.TrimSpace(strings.Split(clientIP, ",")[0])
	}
	clientIP = req.Header.Get("X-Real-IP")
	if clientIP != "" {
		return clientIP
	}
	clientIP = req.Header.Get("http_x_forwarded_for")
	if clientIP != "" {
		return clientIP
	}
	h, _, _ := net.SplitHostPort(req.RemoteAddr)
	return h
}
