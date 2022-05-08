package http

import (
	"backend/app/article"

	fiber "github.com/gofiber/fiber/v2"
)

type Handler struct {
	articleUsecase article.Usecase
}

func New(articleUC article.Usecase) *Handler {
	return &Handler{
		articleUsecase: articleUC,
	}
}

func (h *Handler) Register(app *fiber.App) {
	route := app.Group("/articles")

	//public route
	public := route.Group("")

	public.Post("", h.Create)
	public.Get("", h.Index)

}
