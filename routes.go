package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"notes_program/internal/db"
	"notes_program/internal/notes"
)

func Routes() {
	db, err := db.Open()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	notesService := notes.NewService(db)
	var n notes.Note

	http.HandleFunc("POST /create-note", func(w http.ResponseWriter, r *http.Request) {
		err := json.NewDecoder(r.Body).Decode(&n)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Fatal("deu erro")
		}
		fmt.Println(n)
		notesService.Create("tchauzao")
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}
