package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sunilkkhadka/go-rest-api/internal/models"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {

	payload := struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "Active",
		Message: "The application is running",
		Version: "1.0.0",
	}

	output, err := json.MarshalIndent(payload, "", "\t")
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {

	var movies []models.Movie

	releaseDate, _ := time.Parse("2006-01-02", "1986-03-07")

	highlander := models.Movie{
		ID:          1,
		Title:       "Highlander",
		ReleaseDate: releaseDate,
		MPAARating:  "R",
		RunTime:     116,
		Description: "A very nice movie",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	rotla := models.Movie{
		ID:          1,
		Title:       "Raiders of the lost Arc",
		ReleaseDate: releaseDate,
		MPAARating:  "PG-13",
		RunTime:     126,
		Description: "A Good Movie for kids",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	movies = append(movies, highlander)
	movies = append(movies, rotla)

	output, err := json.MarshalIndent(movies, "", "\t")
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(output)

}
