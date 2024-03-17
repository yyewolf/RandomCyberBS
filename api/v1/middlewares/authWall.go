package middlewares

import (
	"net/http"
	"rcbs/internal/utils"
	"slices"

	"github.com/sirupsen/logrus"
)

func (mr *MiddlewareRessource) AuthWall(roles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims, err := utils.GetToken(r)
			if err != nil {
				logrus.WithError(err).Error("Error getting token from context")
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			// Is true if the user has at least one of the roles
			if !slices.ContainsFunc(claims.Roles, func(role string) bool {
				return slices.Contains(roles, role)
			}) {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
		})
	}
}
