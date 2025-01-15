package notes

import (
	"database/sql"
	"time"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) CreateNewNote(title, body, userId string) error {
	_, err := r.db.Exec(`INSERT INTO note (title, body, created_at, user_id) VALUES ($1, $2, $3, $4)`, title, body, time.Now(), userId)
	if err != nil {
		return err
	}

	return nil
}

// GetAllNotes - получение всех заметок пользователя
func (r *Repo) GetAllNotes(userId string, id int) ([]Note, error) {
	query := `SELECT title, body FROM note WHERE user_id = $1`
	if id == 0 {
		query += ` AND id = $2`
	}
	rows, err := r.db.Query(query, userId, id)
	if err != nil {
		return nil, err
	}

	var notes []Note
	for rows.Next() {
		var note Note
		err = rows.Scan(&note.Title, &note.Body)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, nil
}

// DeleteNote - Удаление заметки
func (r *Repo) DeleteNote(id int, userId string) error {
	_, err := r.db.Exec(`DELETE FROM note WHERE id = $1 AND user_id = $2`, id, userId)
	if err != nil {
		return err
	}

	return nil
}

// UpdateNote - изменение заметки
func (r *Repo) UpdateNote(id int, userId, title, body string) error {
	_, err := r.db.Exec(`UPDATE note SET title = $1, body = $2 WHERE id = $3 AND user_id = $4`, title, body, id, userId)
	if err != nil {
		return err
	}

	return nil
}
