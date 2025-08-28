package sqlite

import (
	"database/sql"
	"errors"

	"github.com/lucasdamasceno96/code/url-shortener/internal/domain"
)

type sqliteRepo struct {
	db *sql.DB
}

func NewSQLiteRepo(db *sql.DB) *sqliteRepo {

	createTableSQL := `CREATE TABLE IF NOT EXISTS short_urls(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		original_url TEXT NOT NULL,
		short_code TEXT NOT NULL UNIQUE,
		created_at TEXT NOT NULL
	);`
	if _, err := db.Exec(createTableSQL); err != nil {
		panic(err)
	}
	return &sqliteRepo{db: db}
}

func (r *sqliteRepo) Save(shortURL *domain.ShortURL) error {
	stmt, err := r.db.Prepare("INSET INTO short_urls(original_url, short_url, shortcode, created_at) VALUES(?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(shortURL.Original, shortURL.ShortCode, shortURL.CreatedAt)
	return err
}

func (r *sqliteRepo) FindByShortCode(shortCode string) (*domain.ShortURL, error) {
	row := r.db.QueryRow("SELECT id, original_url, short_code, created_at FROM short_urls WHERE short_code = ?", shortCode)
	url := &domain.ShortURL{}
	err := row.Scan(&url.ID, &url.Original, &url.ShortCode, &url.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return url, nil
}

func (r *sqliteRepo) Exists(shortCode string) (bool, error) {
	var exists bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM short_urls WHERE short_code =?)", shortCode).Scan(&exists)
	return exists, err
}
