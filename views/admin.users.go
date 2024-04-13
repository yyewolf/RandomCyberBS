package views

import (
	"rcbs/api/v1/controllers/users"
	"rcbs/templa/admin"
	"rcbs/templa/components"

	"github.com/go-fuego/fuego"
)

func adminUsers(c fuego.ContextNoBody) (fuego.Templ, error) {
	searchParams := components.SearchParams{
		Name:    c.QueryParam("name"),
		PerPage: c.QueryParamInt("perPage", 20),
		Page:    c.QueryParamInt("page", 1),
		URL:     "/admin/users",
		Lang:    c.MainLang(),
	}

	r, err := users.GetUsers(searchParams.Name, searchParams.Page, searchParams.PerPage)
	if err != nil {
		return nil, err
	}

	searchParams.Page = r.Page
	searchParams.PerPage = r.PerPage
	searchParams.MaxPage = r.MaxPage

	return admin.UsersPage(r.Users, searchParams), nil
}
