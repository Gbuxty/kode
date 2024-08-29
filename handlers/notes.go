package handlers

import (
	"encoding/json"
	"net/http"
	"kode/models"
	"kode/utils"
)

var notes []models.Note

// GetNotesHandler handles the request to get all notes
func GetNotesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
	utils.LogInfo("Get notes request", map[string]interface{}{
		"method": r.Method,
		"path":   r.URL.Path,
	})
}

// CreateNoteHandler handles the request to create a new note
func CreateNoteHandler(w http.ResponseWriter, r *http.Request) {
	var note models.Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		utils.LogError("Bad request", map[string]interface{}{
			"method": r.Method,
			"path":   r.URL.Path,
			"error":  err.Error(),
		})
		return
	}

	// Validate the note content using Yandex Speller
	if err := validateSpelling(note.Content); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		utils.LogError("Spelling validation failed", map[string]interface{}{
			"method": r.Method,
			"path":   r.URL.Path,
			"error":  err.Error(),
		})
		return
	}

	notes = append(notes, note)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
	utils.LogInfo("Note created", map[string]interface{}{
		"method": r.Method,
		"path":   r.URL.Path,
		"note":   note,
	})
}