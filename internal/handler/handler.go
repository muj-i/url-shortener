package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muj-i/url-shortener/internal/service"
)

type Handler struct {
	service *service.URLService
}

func NewHandler() *Handler {
	return &Handler{service: service.NewURLService()}
}

func (h *Handler) ShortenURL(c *gin.Context) {
	type Request struct {
		URL      string `json:"url"`
		Alias    string `json:"alias"`
		ExpiresIn int    `json:"expires_in"`
	}
	var req Request
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	shortURL, err := h.service.ShortenURL(req.URL, req.Alias, req.ExpiresIn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"short_url": shortURL})
}

func (h *Handler) RedirectURL(c *gin.Context) {
	code := c.Param("code")
	url, err := h.service.GetOriginalURL(code)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}
	c.Redirect(http.StatusMovedPermanently, url)
}

func (h *Handler) GetURLDetails(c *gin.Context) {
	code := c.Param("code")
	details, err := h.service.GetURLDetails(code)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, details)
}

func (h *Handler) DeleteURL(c *gin.Context) {
	code := c.Param("code")
	err := h.service.DeleteURL(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "URL deleted"})
}