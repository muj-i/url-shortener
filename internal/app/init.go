package app

import (
	"github.com/gin-gonic/gin"
	"github.com/muj-i/url-shortener/internal/handler"
)

func Init(r *gin.Engine) {
	h := handler.NewHandler()

	r.POST("/shorten", h.ShortenURL)
	r.GET("/:code", h.RedirectURL)
	r.GET("/info/:code", h.GetURLDetails)
	r.DELETE("/delete/:code", h.DeleteURL)
}
