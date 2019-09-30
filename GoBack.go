package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id          string `json:"Id"`
	Title       string `json:"Title"`
	Description string `json:"description"`
	Content     string `json:"content"`
}

var Articles []Article

func main() {
	Articles = []Article{
		Article{Id: "1", Title: "Hello", Description: "Article description", Content: "Article content"},
		Article{Id: "2", Title: "Hello 2", Description: "Article description", Content: "Article content"},
	}
	handleRequests()
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", http.HandlerFunc(indexHandler))
	myRouter.HandleFunc("/all", http.HandlerFunc(returnAllArticles))
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	index, err := ioutil.ReadFile("index.html")
	if err != nil {
		index = []byte("404 - Your princess is in another castle")
	}
	fmt.Fprintf(w, "%s", index)
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body) //unhandled exception
	var article Article
	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}
