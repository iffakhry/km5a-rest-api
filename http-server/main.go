package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type article struct {
	Id      uint
	Title   string
	Content string
}

type book struct {
	Id     uint
	Title  string
	Author string
}

// dummy data
var data = []article{
	article{1, "lorem", "lorem"},
	article{2, "ipsum", "ipsum"},
	article{3, "abcd", "abcd"},
}

var dataBooks = []book{
	{1, "One Piece 1", "Oda"},
	{2, "One Piece 2", "Oda"},
	{3, "One Piece 3", "Oda"},
	{4, "One Piece 4", "Oda"},
}

func main() {
	// POST /reservations
	//define endpoint
	http.HandleFunc("/articles", articlesController)
	http.HandleFunc("/books", booksController)
	fmt.Println("starting web server at http://localhost:8080/")
	// start service/server at port 8080
	http.ListenAndServe(":8080", nil)
}

func articlesController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		var result, err = json.Marshal(data)
		err = errors.New("data tidak ada")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	} else if r.Method == "POST" {
		var responseData = map[string]any{
			"message": "success insert data",
			"status":  true,
		}
		var result, err = json.Marshal(responseData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func booksController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		var result, err = json.Marshal(dataBooks)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}
