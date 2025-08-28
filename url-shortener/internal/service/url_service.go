package service

import (
	"errors"
	"log"
	"time"

	"github.com/lucasdamasceno96/code/url-shortener/internal/config"
	"github.com/lucasdamasceno96/code/url-shortener/internal/domain"
	"github.com/lucasdamasceno96/code/url-shortener/internal/repository"
	"github.com/lucasdamasceno96/code/url-shortener/internal/util"
)

type URLService interface {
	ShortenURL(originalURL string) (string, error)
	Redirect(shortCode string) (string, error)
}

type urlService struct {
	repo   repository.URLRepository
	config config.Config
}

func NewURLService(r repository.URLRepository, c config.Config) URLService {
	return &urlService{repo: r, config: c}
}

func (s *urlService) ShortenURL(originalURL string) (string, error) {
	var newShortCode string
	var exists bool
	var err error

	for {
		newShortCode = util.GenerateShortCode(s.config.FixedLetters)
		exists, err = s.repo.Exists(newShortCode)
		if err != nil {
			return "", err
		}
		if !exists {
			break
		}
	}

	shortURL := &domain.ShortURL{
		Original:  originalURL,
		ShortCode: newShortCode,
		CreatedAt: time.Now().UTC().Format(time.RFC3339),
	}

	if err := s.repo.Save(shortURL); err != nil {
		log.Printf("ERROR: Service layer failed to save at repo: %v", err)
		return "", err
	}

	return s.config.BaseURL + "/" + newShortCode, nil
}

func (s *urlService) Redirect(shortCode string) (string, error) {
	shortURL, err := s.repo.FindByShortCode(shortCode)
	if err != nil {
		return "", err
	}
	if shortURL == nil {
		return "", errors.New("short code not found")
	}
	return shortURL.Original, nil
}
