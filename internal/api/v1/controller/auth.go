package controller

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	internalerrors "github.com/JulioZittei/go-job-mail-service/internal/domain/internalErrors"
	oidc "github.com/coreos/go-oidc/v3/oidc"
	"github.com/go-chi/render"
	"github.com/golang-jwt/jwt/v5"
)

type EmailKey string

type MyCustomClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

const (
	EMAIL_KEY EmailKey = "email"
	EMAIL     string   = "email"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get("Authorization")
		if authToken == "" {
			render.Status(r, 401)
			render.JSON(w, r, &internalerrors.ErrorResponse{
				Code:          http.StatusUnauthorized,
				Status:        http.StatusText(http.StatusUnauthorized),
				Title:         "Unauthorized",
				Detail:        "Missing header 'Authorization'",
				Instance:      r.RequestURI,
				InvalidParams: []internalerrors.ErrorsParam{},
			})
			return
		}

		authToken = strings.Replace(authToken, "Bearer ", "", 1)
		provider, err := oidc.NewProvider(r.Context(), os.Getenv("KEYCLOAK_URL"))

		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			err := internalerrors.NewErrInternal()
			render.JSON(w, r, &internalerrors.ErrorResponse{
				Code:          err.StatusCode,
				Status:        http.StatusText(err.StatusCode),
				Title:         err.Title,
				Detail:        "Error connecting to the provider",
				Instance:      r.RequestURI,
				InvalidParams: []internalerrors.ErrorsParam{},
			})
			return
		}

		verifier := provider.Verifier(&oidc.Config{ClientID: "mail-service-user"})
		_, err = verifier.Verify(r.Context(), authToken)
		if err != nil {
			render.Status(r, 401)
			render.JSON(w, r, &internalerrors.ErrorResponse{
				Code:          http.StatusUnauthorized,
				Status:        http.StatusText(http.StatusUnauthorized),
				Title:         "Unauthorized",
				Detail:        "Invalid authToken",
				Instance:      r.RequestURI,
				InvalidParams: []internalerrors.ErrorsParam{},
			})
			return
		}

		claims, err := getClaims(authToken)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			err := internalerrors.NewErrInternal()
			render.JSON(w, r, &internalerrors.ErrorResponse{
				Code:          err.StatusCode,
				Status:        http.StatusText(err.StatusCode),
				Title:         err.Title,
				Detail:        "Error getting claims",
				Instance:      r.RequestURI,
				InvalidParams: []internalerrors.ErrorsParam{},
			})
			return
		}

		email := claims[EMAIL]
		ctx := context.WithValue(r.Context(), EMAIL_KEY, email)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getClaims(authToken string) (jwt.MapClaims, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(authToken, jwt.MapClaims{})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("token inválido")
	}

	return claims, nil
}
