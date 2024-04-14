package views

import (
	"github.com/go-fuego/fuego"
)

func Routes(s *fuego.Server) {
	// Public Pages
	fuego.All(s, "/login", login)

	staticRoutes := fuego.Group(s, "/static")
	fuego.GetStd(staticRoutes, "/robots.txt", robots)
	fuego.GetStd(staticRoutes, "/tailwind.min.css", tailwind)

	// Admin Pages
	adminRoutes := fuego.Group(s, "/admin")
	fuego.Get(adminRoutes, "/users", adminUsers)
}
