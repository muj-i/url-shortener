package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/muj-i/url-shortener/internal/model"
	"github.com/muj-i/url-shortener/internal/repository"
)

type URLService struct {
	repo *repository.URLRepository
}

func NewURLService() *URLService {
	return &URLService{repo: repository.NewURLRepository()}
}

func (s *URLService) ShortenURL(longURL string, alias string, expiresIn int) (string, error) {
	if alias == "" {
		alias = uuid.New().String()[:6]
	}
	_, err := s.repo.FindByCode(alias)
	if err == nil {
		return "", errors.New("alias already exists")
	}
	var expiry *time.Time
	if expiresIn > 0 {
		t := time.Now().Add(time.Duration(expiresIn) * time.Second)
		expiry = &t
	}
	url := &model.URL{
		Code:      alias,
		LongURL:   longURL,
		CreatedAt: time.Now(),
		ExpiresAt: expiry,
	}
	return s.repo.Save(url)
}

func (s *URLService) GetOriginalURL(code string) (string, error) {
	url, err := s.repo.FindByCode(code)
	if err != nil || (url.ExpiresAt != nil && url.ExpiresAt.Before(time.Now())) {
		return "", errors.New("not found or expired")
	}
	s.repo.IncrementClick(code)
	return url.LongURL, nil
}

func (s *URLService) GetURLDetails(code string) (*model.URL, error) {
	return s.repo.FindByCode(code)
}

func (s *URLService) DeleteURL(code string) error {
	return s.repo.DeleteByCode(code)
}