package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/shiba-hiro/note-api-by-go/repository"
)

func TestCreateNote(t *testing.T) {
	e := echo.New()

	db, err := repository.OpenDbConnection()
	if err != nil {
		e.Logger.Fatalf("Cannot open Database: %v\n", db)
		return
	}
	defer db.Close()

	noteJSON := `{"title":"My First Note","content":"I started to take note."}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/notes", strings.NewReader(noteJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, createNote(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestGetNote(t *testing.T) {
	e := echo.New()

	db, err := repository.OpenDbConnection()
	if err != nil {
		e.Logger.Fatalf("Cannot open Database: %v\n", db)
		return
	}
	defer db.Close()

	id := uuid.New().String()
	if db.Create(&repository.Note{ID: id, Title: "foo", Content: "bar", CreatedAt: time.Now(), UpdatedAt: time.Now()}).Error != nil {
		e.Logger.Fatalf("Cannot store test data: %v\n", db)
		return
	}

	req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/v1/notes/%s", id), nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, showNote(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestUpdateNote(t *testing.T) {
	e := echo.New()

	db, err := repository.OpenDbConnection()
	if err != nil {
		e.Logger.Fatalf("Cannot open Database: %v\n", db)
		return
	}
	defer db.Close()

	id := uuid.New().String()
	if db.Create(&repository.Note{ID: id, Title: "foo", Content: "bar", CreatedAt: time.Now(), UpdatedAt: time.Now()}).Error != nil {
		e.Logger.Fatalf("Cannot store test data: %v\n", db)
		return
	}

	noteJSON := `{"title":"My First Note","content":"I started to take note."}`

	req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/v1/notes/%s", id), strings.NewReader(noteJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, updateNote(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestDeleteNote(t *testing.T) {
	e := echo.New()

	db, err := repository.OpenDbConnection()
	if err != nil {
		e.Logger.Fatalf("Cannot open Database: %v\n", db)
		return
	}
	defer db.Close()

	id := uuid.New().String()
	if db.Create(&repository.Note{ID: id, Title: "foo", Content: "bar", CreatedAt: time.Now(), UpdatedAt: time.Now()}).Error != nil {
		e.Logger.Fatalf("Cannot store test data: %v\n", db)
		return
	}

	req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/v1/notes/%s", id), nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, deleteNote(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
	}
}
