package domain

type ShortURL struct {
	ID        int64  `json:"id"`
	Original  string `json:"original_url"`
	ShortCode string `json:"short_code"`
	CreatedAt string `json:"created_at"`
}
