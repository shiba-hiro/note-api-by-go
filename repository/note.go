package repository

import (
	"time"

	"github.com/google/uuid"
)

type Note struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ListNotes() ([]Note, error) {
	// TODO: Implement
	return []Note{
		{
			ID:        "6eca4cf8-91f9-11ea-96cf-00163ef6bb2e",
			Title:     "foo",
			Content:   "bar",
			CreatedAt: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
	}, nil
}

func GetNote(id string) (Note, error) {
	// TODO: Implement
	return Note{
		ID:        id,
		Title:     "foo",
		Content:   "bar",
		CreatedAt: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
	}, nil
}

func CreateNote(title string, content string) (Note, error) {
	// TODO: Implement
	return Note{
		ID:        uuid.New().String(),
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func UpdateNote(id string, title string, content string) (Note, error) {
	// TODO: Implement
	return Note{
		ID:        id,
		Title:     title,
		Content:   content,
		CreatedAt: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Now(),
	}, nil
}

func DeleteNote(id string) error {
	// TODO: Implement
	return nil
}
