package main

import (
	"crypto/rand"
	"embed"
	"html/template"
	"log"
	"net/http"
	"strings"
)

//go:embed templates/*.html static/*
var content embed.FS

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
	tmpl, err := template.ParseFS(content, "templates/index.html")
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

	// Menyajikan file statis (misalnya CSS) dari embed
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(content))))

	// Jalankan server di port 8011
	log.Println("Server started at http://0.0.0.0:8011")
	err := http.ListenAndServe("0.0.0.0:8011", nil)
	if err != nil {
		log.Fatal(err)
	}
}
