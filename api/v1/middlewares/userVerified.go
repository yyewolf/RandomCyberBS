package middlewares

import (
	"net/http"
	"rcbs/internal/utils"

	"github.com/sirupsen/logrus"
)

func IsVerified(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := utils.GetUser(r)
		if err != nil {
			logrus.WithError(err).Error("Error getting token from context")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if !user.Verified {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
