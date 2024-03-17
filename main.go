package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Open DB
	db, err := sql.Open("sqlite3", "./urls.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	//TODO: Write post function
	r.Post("/create", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Created"))
	})
	//TODO: Write get function
	r.Get("/{url}", func(w http.ResponseWriter, r *http.Request) {
		url := chi.URLParam(r, "url")
		var redirect string
		err = db.QueryRow("select redirect from urls where from = ?", url).Scan(&redirect)
		if err != nil {
			log.Fatal(err)
		}

		if redirect == "" {
			http.Redirect(w, r, "https://vt.edu", http.StatusTemporaryRedirect)
		}

		http.Redirect(w, r, redirect, http.StatusPermanentRedirect)
	})

	http.ListenAndServe(":3000", r)
}
