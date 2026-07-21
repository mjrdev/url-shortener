package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/mjrdev/internal/handler"
	"github.com/mjrdev/internal/middleware"
)

func Router(router chi.Router) {
	router.Route("/api", func(r chi.Router) {
		r.Post("/auth", handler.Authenticate)
		r.Post("/user", handler.UserCreate)

		r.Group(func(r chi.Router) {
			r.Use(middleware.JwtMiddleware)

			r.Get("/me", handler.UserShow)

			r.Post("/url", handler.StoreUrl)
			r.Delete("/url/{id}", handler.DeleteUrl)
			r.Get("/url/{short_url}", handler.ShowUrl)
			r.Get("/url", handler.ListUrl)
		})
	})
	router.Get("/u/{path}", handler.Redirect)
}
