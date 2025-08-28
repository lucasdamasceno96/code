// internal/service/url_service_test.go
package service

import (
	"testing"

	"github.com/lucasdamasceno96/code/url-shortener/internal/config"
	"github.com/lucasdamasceno96/code/url-shortener/internal/domain"
)

// mockRepository is our mock implementation of the real repository.
type mockRepository struct {
	// We use these fields to control the mock's behavior in each test.
	SaveFunc            func(shortURL *domain.ShortURL) error
	ExistsFunc          func(shortCode string) (bool, error)
	FindByShortCodeFunc func(shortCode string) (*domain.ShortURL, error)
}

// We implement the URLRepository interface methods.
func (m *mockRepository) Save(shortURL *domain.ShortURL) error {
	return m.SaveFunc(shortURL)
}

func (m *mockRepository) Exists(shortCode string) (bool, error) {
	return m.ExistsFunc(shortCode)
}
func (m *mockRepository) FindByShortCode(shortCode string) (*domain.ShortURL, error) {
	return m.FindByShortCodeFunc(shortCode)
}

// Now, the test for the ShortenURL function of our service.
func TestURLService_ShortenURL(t *testing.T) {
	// Arrange: Set up the test environment.

	// 1. Create our mock.
	mockRepo := &mockRepository{}

	// 2. Define a configuration for the test.
	testConfig := config.Config{
		BaseURL:      "http://test.com",
		FixedLetters: "tst-",
	}

	// 3. Create an instance of the service, injecting our mock and test config.
	urlSvc := NewURLService(mockRepo, testConfig)

	// 4. Define the mock's expected behavior for this test.
	// When the service calls 'Exists', the mock should return 'false' (code does not exist).
	mockRepo.ExistsFunc = func(shortCode string) (bool, error) {
		return false, nil
	}
	// When the service calls 'Save', it should not return any error.
	mockRepo.SaveFunc = func(shortURL *domain.ShortURL) error {
		return nil
	}

	// Act: Execute the function we want to test.
	shortURL, err := urlSvc.ShortenURL("https://example.com")

	// Assert: Check if the result is as expected.
	if err != nil {
		t.Fatalf("ShortenURL returned an unexpected error: %v", err)
	}

	if shortURL == "" {
		t.Errorf("ShortenURL returned an empty URL")
	}

	// We check if the returned URL contains the base from our test config.
	expectedBase := "http://test.com/tst-"
	if len(shortURL) != len(expectedBase)+5 {
		t.Errorf("The short URL has an unexpected format: %s", shortURL)
	}
}
