package goaext

import (
	"compress/flate"
	"errors"
	"net/http"

	"github.com/ldd27/go-lib/zaplog"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/goadesign/goa/middleware/gzip"
)

type JwtConfig struct {
	Secret        string
	Scheme        *goa.JWTSecurity
	UseMiddleware func(service *goa.Service, middleware goa.Middleware)
}

type LogRequestConfig struct {
	Verbose bool
}

type Option struct {
	Debug       bool
	ServiceName string
	Log         *zaplog.ExtendLogger

	// middleware
	EnableRequestID    bool
	EnableErrorHandler bool
	EnableRecover      bool
	EnableGZIP         bool
	EnableRealIP       bool
	EnableCORS         bool
	EnableJWT          bool
	JWTConf            JwtConfig
	EnableLogRequest   bool
	LogRequestConf     LogRequestConfig
	InternalError      error
}

var defaultGOAConf = Option{
	EnableRequestID:    true,
	EnableErrorHandler: true,
	EnableRecover:      true,
	EnableGZIP:         true,
	EnableRealIP:       true,
	EnableCORS:         true,
	EnableJWT:          true,
	EnableLogRequest:   true,
	JWTConf:            JwtConfig{},
	LogRequestConf:     LogRequestConfig{},
	InternalError:      errors.New(http.StatusText(http.StatusInternalServerError)),
}

func NewGOA(opts ...func(*Option)) (*goa.Service, error) {
	conf := defaultGOAConf
	for _, o := range opts {
		o(&conf)
	}

	service := goa.New(conf.ServiceName)

	if conf.Log != nil {
		service.WithLogger(conf.Log)
	}

	// middleware
	if conf.EnableRequestID {
		service.Use(middleware.RequestID())
	}

	if conf.EnableErrorHandler {
		service.Use(ErrorHandler(conf.InternalError, service, conf.Debug))
	}

	if conf.EnableRecover {
		service.Use(middleware.Recover())
	}

	if conf.EnableGZIP {
		service.Use(gzip.Middleware(flate.BestCompression))
	}

	if conf.EnableRealIP {
		service.Use(NewRealIPMiddleware())
	}

	if conf.EnableCORS {
		service.Use(NewCORSMiddleware())
	}

	if conf.EnableLogRequest {
		service.Use(LogRequest(conf.LogRequestConf.Verbose))
	}

	if conf.EnableJWT {
		jwtMiddleware, err := NewJWTMiddleware(conf.JWTConf.Secret, conf.JWTConf.Scheme)
		if err != nil {
			return nil, err
		}
		conf.JWTConf.UseMiddleware(service, jwtMiddleware)
	}

	return service, nil
}
