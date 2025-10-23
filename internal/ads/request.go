package ads

type CreateAdRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
	UserID      string `json:"user_id"`
}
