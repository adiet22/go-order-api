package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/adiet95/go-order-api/src/libs"
)

func CheckAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		headerToken := r.Header.Get("Authorization")

		if !strings.Contains(headerToken, "Bearer") {
			libs.New("invalid header type", 401, true).Send(w)
			return
		}
		token := strings.Replace(headerToken, "Bearer ", "", -1)

		checkToken, err := libs.CheckToken(token)
		if err != nil {
			libs.New(err.Error(), 401, true).Send(w)
			return
		}
		ctx := context.WithValue(r.Context(), "email", checkToken.Email)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
