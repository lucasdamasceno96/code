package repository

import "github.com/lucasdamasceno96/code/url-shortener/internal/domain"

type URLRepository interface {
	Save(shortURL *domain.ShortURL) error
	FindByShortCode(shortCode string) (*domain.ShortURL, error)
	Exists(shortCode string) (bool, error)
}
