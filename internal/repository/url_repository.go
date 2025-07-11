package repository

import (
	"errors"
	"sync"

	"github.com/muj-i/url-shortener/internal/model"
)

var (
	mockDB = make(map[string]*model.URL)
	mu     sync.Mutex
)

type URLRepository struct{}

func NewURLRepository() *URLRepository {
	return &URLRepository{}
}

func (r *URLRepository) Save(url *model.URL) (string, error) {
	mu.Lock()
	defer mu.Unlock()
	mockDB[url.Code] = url
	return "http://localhost:8080/" + url.Code, nil
}

func (r *URLRepository) FindByCode(code string) (*model.URL, error) {
	mu.Lock()
	defer mu.Unlock()
	if url, ok := mockDB[code]; ok {
		return url, nil
	}
	return nil, errors.New("not found")
}

func (r *URLRepository) DeleteByCode(code string) error {
	mu.Lock()
	defer mu.Unlock()
	delete(mockDB, code)
	return nil
}

func (r *URLRepository) IncrementClick(code string) {
	mu.Lock()
	defer mu.Unlock()
	if url, ok := mockDB[code]; ok {
		url.ClickCount++
	}
}