package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/anastasiia-shurkina-axon/go-first/models"
	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "homePage endpoint")
}

func (s *server) articleList(w http.ResponseWriter, r *http.Request) {
	res, err := s.articleService.List()
	if err != nil {
		log.Printf("internal error: %v", err)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func (s *server) articleDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	id, err := strconv.Atoi(key)
	if err != nil {
		log.Printf("bad request error: %v", err)
		return
	}

	res, err := s.articleService.Read(id)
	if err != nil {
		log.Printf("internal error: %v", err)
		return
	}
	json.NewEncoder(w).Encode(res)

}

func (s *server) createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var article models.Article
	err := json.Unmarshal(reqBody, &article)
	if err != nil {
		log.Printf("bad request error %v", err)
		return
	}

	res, err := s.articleService.Create(&article)

	if err != nil {
		log.Printf("internal error: %v", err)
		return
	}

	json.NewEncoder(w).Encode(res)
}
