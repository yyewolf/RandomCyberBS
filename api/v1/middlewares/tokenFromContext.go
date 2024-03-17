package middlewares

import (
	"context"
	"encoding/json"
	"net/http"
	"rcbs/internal/auth"
	"rcbs/internal/utils"
	"rcbs/models"
	"slices"

	"github.com/go-fuego/fuego"
	"github.com/golang-jwt/jwt/v5"
)

func (mr *MiddlewareRessource) TokenFromContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the authorizationHeader from the header
		token := fuego.TokenFromCookie(r)
		if token == "" {
			// Unauthenticated, might be legit
			next.ServeHTTP(w, r)
			return
		}

		// Validate the token
		t, err := mr.s.Security.ValidateToken(token)
		if err != nil {
			// Remove the cookie fuego.JWTCookieName
			http.SetCookie(w, &http.Cookie{
				Name:   fuego.JWTCookieName,
				Value:  "",
				MaxAge: -1,
			})

			next.ServeHTTP(w, r)
			return
		}

		if !t.Valid {
			// Remove the cookie fuego.JWTCookieName
			http.SetCookie(w, &http.Cookie{
				Name:   fuego.JWTCookieName,
				Value:  "",
				MaxAge: -1,
			})

			next.ServeHTTP(w, r)
			return
		}

		// Infernal hell,
		// TODO: Find a better way to do this
		m := t.Claims.(jwt.MapClaims)
		d, _ := json.Marshal(m)
		var authClaims auth.AuthClaims
		json.Unmarshal(d, &authClaims)

		// Get user from DB and update cookie if necessary
		u, err := models.Db.Users.FindById(authClaims.ID)
		if err != nil {
			// Remove the cookie fuego.JWTCookieName
			http.SetCookie(w, &http.Cookie{
				Name:   fuego.JWTCookieName,
				Value:  "",
				MaxAge: -1,
			})

			next.ServeHTTP(w, r)
			return
		}

		// Update the cookie if necessary
		if !slices.Equal(u.Roles, authClaims.Roles) {
			authClaims.Roles = u.Roles

			mr.s.Security.GenerateTokenToCookies(authClaims, w)
		}

		// Set the subject and roles in the context
		ctx := r.Context()
		ctx = context.WithValue(ctx, utils.TokenKey, authClaims)
		ctx = context.WithValue(ctx, utils.UserKey, u)
		r = r.WithContext(ctx)

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
