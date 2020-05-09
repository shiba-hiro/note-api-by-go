package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/shiba-hiro/note-api-by-go/repository"
)

type NoteRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func Register(group *echo.Group) {
	notes := group.Group("/notes")
	notes.GET("", listNotes)
	notes.GET("/:id", showNote)
	notes.POST("", createNote)
	notes.PUT("/:id", updateNote)
	notes.DELETE("/:id", deleteNote)
}

func listNotes(c echo.Context) error {
	c.Logger().Info("listNotes called")
	notes, err := repository.ListNotes()
	if err != nil {
		message := fmt.Sprintf("Unexpected error occurred: %v", err)
		c.Logger().Error(message)
		return c.String(http.StatusInternalServerError, fmt.Sprintf("{\"message\": %s}", message))
	}
	c.Logger().Info("Successfully fetched notes")
	return c.JSON(http.StatusOK, &notes)
}

func showNote(c echo.Context) error {
	id := c.Param("id")
	c.Logger().Infof("showNote called. id: %s", id)
	note, err := repository.GetNote(id)
	if err != nil {
		message := fmt.Sprintf("Unexpected error occurred: %v", err)
		c.Logger().Error(message)
		return c.String(http.StatusInternalServerError, fmt.Sprintf("{\"message\": %s}", message))
	}
	c.Logger().Info("Successfully fetched note")
	return c.JSON(http.StatusOK, &note)
}

func createNote(c echo.Context) error {
	c.Logger().Info("createNote called")

	r := new(NoteRequest)
	if err := c.Bind(r); err != nil {
		message := fmt.Sprintf("Cannot parse the request: %v", err)
		c.Logger().Warn(message)
		return c.String(http.StatusBadRequest, fmt.Sprintf("{\"message\": %s}", message))
	}

	note, err := repository.CreateNote(r.Title, r.Content)
	if err != nil {
		message := fmt.Sprintf("Unexpected error occurred: %v", err)
		c.Logger().Error(message)
		return c.String(http.StatusInternalServerError, fmt.Sprintf("{\"message\": %s}", message))
	}
	c.Logger().Info("Successfully created note")
	return c.JSON(http.StatusOK, &note)
}

func updateNote(c echo.Context) error {
	id := c.Param("id")
	c.Logger().Infof("updateNote called. id: %s", id)

	r := new(NoteRequest)
	if err := c.Bind(r); err != nil {
		message := fmt.Sprintf("Cannot parse the request: %v", err)
		c.Logger().Warn(message)
		return c.String(http.StatusBadRequest, fmt.Sprintf("{\"message\": %s}", message))
	}

	note, err := repository.UpdateNote(id, r.Title, r.Content)
	if err != nil {
		message := fmt.Sprintf("Unexpected error occurred: %v", err)
		c.Logger().Error(message)
		return c.String(http.StatusInternalServerError, fmt.Sprintf("{\"message\": %s}", message))
	}
	c.Logger().Info("Successfully updated note")
	return c.JSON(http.StatusOK, &note)
}

func deleteNote(c echo.Context) error {
	id := c.Param("id")
	c.Logger().Infof("deleteNote called. id: %s", id)
	err := repository.DeleteNote(id)
	if err != nil {
		message := fmt.Sprintf("Unexpected error occurred: %v", err)
		c.Logger().Error(message)
		return c.String(http.StatusInternalServerError, fmt.Sprintf("{\"message\": %s}", message))
	}
	c.Logger().Info("Successfully deleted note")
	return c.NoContent(http.StatusNoContent)
}
