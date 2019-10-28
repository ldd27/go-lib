package goaext

import (
	"context"
	"net/http"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware/security/jwt"
)

func NewJWTMiddleware(secret string, scheme *goa.JWTSecurity) (goa.Middleware, error) {
	return jwt.New(secret, jwtTokenValidation(), scheme), nil
}

func jwtTokenValidation() goa.Middleware {
	f := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		token := jwt.ContextJWT(ctx)
		claims, ok := token.Claims.(jwtgo.MapClaims)
		if !ok {
			return jwt.ErrJWTError("unsupported claims shape")
		}

		if val, ok := claims["sub"].(string); !ok || val == "" {
			return jwt.ErrJWTError("invalid token")
		}

		return nil
	}

	validationHandler, _ := goa.NewMiddleware(f)
	return validationHandler
}
