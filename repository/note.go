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
	notes := []Note{}
	err := DB.Find(&notes).Error
	return notes, err
}

func GetNote(id string) (Note, error) {
	note := Note{ID: id}
	err := DB.Where(&note).First(&note).Error
	return note, err
}

func CreateNote(title string, content string) (Note, error) {
	note := Note{
		ID:        uuid.New().String(),
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := DB.Create(&note).Error
	return note, err
}

func UpdateNote(id string, title string, content string) (Note, error) {
	note := Note{ID: id}
	if err := DB.Where(&note).FirstOrCreate(&note).Error; err != nil {
		return note, err
	}

	note.Title = title
	note.Content = content
	note.UpdatedAt = time.Now()
	err := DB.Save(&note).Error
	return note, err
}

func DeleteNote(id string) error {
	return DB.Delete(&Note{ID: id}).Error
}
