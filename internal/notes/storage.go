package notes

import "time"

type Note struct {
	Id        int        `json:"id,omitempty"`
	UserId    string     `json:"user_id,omitempty"`
	Title     string     `json:"title"`
	Body      string     `json:"body"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}
