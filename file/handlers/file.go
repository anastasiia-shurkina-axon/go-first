package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "homePage endpoint")
}

func (s *server) fileList(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("file list")
}

func (s *server) fileDetails(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("file details")
}

func (s *server) createNewFile(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("new created file")
}

func (s *server) deletefile(w http.ResponseWriter, r *http.Request) {
}
