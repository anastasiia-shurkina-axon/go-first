package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
)

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	Articles  = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Hello World", Content: "Hello, my dear Ukraine. I like coding. I like cats"},
		Article{Id: "2", Title: "Goodbye", Desc: "Goodbye World", Content: "I like dogs. London is a capital of Great Britain. Goodbye!"},
	}
	handleRequests()
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage).Methods("GET")

	myRouter.HandleFunc("/articles/{id}", articleDetails).Methods("GET")
	myRouter.HandleFunc("/articles", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/articles", articleList).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "homePage endpoint")
}

func articleList(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Articles)
}

func articleDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Fprintf(w, "Key: " + key)

	for _, article := range Articles {
        if article.Id == key {
            json.NewEncoder(w).Encode(article)
        }
    }
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(reqBody))

	var article Article
	json.Unmarshal(reqBody, &article)

	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(Articles)
}

type Article struct {
	Id string `json:id`
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article
