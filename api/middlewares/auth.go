package middlewares

import (
	"context"
	"errors"
	"myapi/apperrors"
	"myapi/common"
	"net/http"
	"strings"

	"google.golang.org/api/idtoken"
)

const (
	googleClientID = "695376933230-0glkcevcuulcjidqb06haadbu0i6htk8.apps.googleusercontent.com"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorization := req.Header.Get("Authorization")
		authHeaders := strings.Split(authorization, " ")
		if len(authHeaders) != 2 {
			err := apperrors.RequiredAuthorizationHeader.Wrap(errors.New("invalid req header"), "invalid header")
			apperrors.ErrorHandler(w, req, err)
			return
		}
		bearer, idToken := authHeaders[0], authHeaders[1]
		if bearer != "Bearer" || idToken == "" {
			err := apperrors.RequiredAuthorizationHeader.Wrap(errors.New("invalid req header"), "invalid header")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		tokenValidator, err := idtoken.NewValidator(context.Background())
		if err != nil {
			err = apperrors.CannotMakeValidator.Wrap(err, "internal auth error")
			apperrors.ErrorHandler(w, req, err)
			return
		}
		payload, err := tokenValidator.Validate(context.Background(), idToken, googleClientID)
		if err != nil {
			err = apperrors.Unauthorizated.Wrap(err, "invalid id token")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		name, ok := payload.Claims["name"]
		if !ok {
			err = apperrors.Unauthorizated.Wrap(err, "invalid id token")
			apperrors.ErrorHandler(w, req, err)
			return
		}
		req = common.SetuserName(req, name.(string))

		// トークンが正しい場合は次のハンドラに進む
		next.ServeHTTP(w, req)
	})
}
