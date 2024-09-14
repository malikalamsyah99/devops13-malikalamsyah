package main

import (
	"crypto/rand"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type PasswordData struct {
	Password string
}

// GenerateRandomPassword generates a random password based on the ISO 27001 recommendation
func GenerateRandomPassword(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"0123456789" +
		"!@#$%^&*()_-+={}[]<>?"

	var password strings.Builder
	buffer := make([]byte, length)
	if _, err := rand.Read(buffer); err != nil {
		return "", err
	}

	for _, b := range buffer {
		password.WriteByte(charset[b%byte(len(charset))])
	}

	return password.String(), nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {
		password, err := GenerateRandomPassword(16) // Panjang password bisa disesuaikan
		if err != nil {
			http.Error(w, "Error generating password", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, PasswordData{Password: password})
	} else {
		tmpl.Execute(w, nil)
	}
}

func main() {
	http.HandleFunc("/", indexHandler)

	// Menyajikan file statis seperti CSS
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))))

	// Jalankan server di port 8080
	log.Println("Server started at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
