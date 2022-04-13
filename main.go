package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type ArticleHandler struct {
	db *sql.DB
}

type Article struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/go-first")

	if err != nil {
		panic(err)
	}

	handleRequests(db)
}

func handleRequests(db *sql.DB) {

	ah := &ArticleHandler{db}
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage).Methods("GET")

	myRouter.HandleFunc("/articles/{id}", ah.articleDetails).Methods("GET")
	myRouter.HandleFunc("/articles", ah.createNewArticle).Methods("POST")
	myRouter.HandleFunc("/articles", ah.articleList).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "homePage endpoint")
}

func (h *ArticleHandler) articleList(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	rows, err := h.db.Query("select id, title, description, content from articles")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var articles []*Article
	for rows.Next() {
		art := &Article{}
		err := rows.Scan(&art.Id, &art.Title, &art.Desc, &art.Content)
		if err != nil {
			panic(err)
		}
		articles = append(articles, art)
	}

	json.NewEncoder(w).Encode(articles)
}

func (h *ArticleHandler) articleDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Println("Key: " + key)

	rows, err := h.db.Query("select id, title, description, content from articles where id = ?", key)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	article := &Article{}
	for rows.Next() {
		err := rows.Scan(&article.Id, &article.Title, &article.Desc, &article.Content)
		if err != nil {
			panic(err)
		}
	}

	json.NewEncoder(w).Encode(article)
}

func (h *ArticleHandler) createNewArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	reqBody, _ := ioutil.ReadAll(r.Body)

	var article Article
	json.Unmarshal(reqBody, &article)

	stmt, err := h.db.Prepare("insert into articles(title, description, content) values (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(article.Title, article.Desc, article.Content)
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	article.Id = int(lastId)

	json.NewEncoder(w).Encode(article)
}
