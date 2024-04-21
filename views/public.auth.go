package views

import (
	"rcbs/templa/public"

	"github.com/go-fuego/fuego"
)

func login(c fuego.ContextNoBody) (fuego.Templ, error) {
	return public.LoginPage(), nil
}

func register(c fuego.ContextNoBody) (fuego.Templ, error) {
	return public.RegisterPage(), nil
}
