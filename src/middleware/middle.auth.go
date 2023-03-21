package middleware

import (
	"context"
	"lectronic/src/libs"
	"net/http"
	"strings"
)

type UserID string
type Role string

func AuthMiddle(role ...string) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.Header().Set("Content-type", "application/json")

			var header string
			var valid bool

			if header = r.Header.Get("Authorization"); header == "" {
				libs.GetResponse("You need to login first", 401, true).Send(w)
				return
			}

			if !strings.Contains(header, "Bearer") {
				libs.GetResponse("Invalid header", 401, true).Send(w)
				return
			}

			tokens := strings.Replace(header, "Bearer ", "", -1)

			checkToken, err := libs.CheckToken(tokens)
			if err != nil {
				libs.GetResponse(err.Error(), 401, true).Send(w)
				return
			}

			for _, rl := range role {
				if rl == checkToken.Role {
					valid = true
				}
			}

			if !valid {
				libs.GetResponse("You do not have permission", 401, true).Send(w)
				return
			}

			ctx := context.WithValue(r.Context(), "user", checkToken.UserID)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
