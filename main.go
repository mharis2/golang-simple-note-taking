package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

type Note struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var notes []Note
var idCounter int
var noteMutex sync.Mutex

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/notes", getNotes).Methods("GET")
	router.HandleFunc("/notes/{id}", getNote).Methods("GET")
	router.HandleFunc("/notes", createNote).Methods("POST")
	router.HandleFunc("/notes/{id}", updateNote).Methods("PUT")
	router.HandleFunc("/notes/{id}", deleteNote).Methods("DELETE")
	fmt.Println("Server is running on port 8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

func getNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for _, note := range notes {
		if note.ID == id {
			json.NewEncoder(w).Encode(note)
			return
		}
	}
	http.Error(w, "Note not found", http.StatusNotFound)
}

func createNote(w http.ResponseWriter, r *http.Request) {
	var newNote Note
	json.NewDecoder(r.Body).Decode(&newNote)

	noteMutex.Lock()
	newNote.ID = idCounter
	idCounter++
	noteMutex.Unlock()

	notes = append(notes, newNote)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newNote)
}

func updateNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var updatedNote Note
	json.NewDecoder(r.Body).Decode(&updatedNote)

	for i, note := range notes {
		if note.ID == id {
			notes[i].Title = updatedNote.Title
			notes[i].Content = updatedNote.Content
			json.NewEncoder(w).Encode(notes[i])
			return
		}
	}
	http.Error(w, "Note not found", http.StatusNotFound)
}

func deleteNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, note := range notes {
		if note.ID == id {
			notes = append(notes[:i], notes[i+1:]...)
			json.NewEncoder(w).Encode(note)
			return
		}
	}
	http.Error(w, "Note not found", http.StatusNotFound)
}
