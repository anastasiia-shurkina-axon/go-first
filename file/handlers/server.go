package handlers

import (
	"database/sql"
	"github.com/anastasiia-shurkina-axon/go-first/file/middlewares"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"

	"github.com/go-chi/chi/v5/middleware"
)

type Server interface {
	GetRouter() *chi.Mux
}
type server struct{}

func NewServer() Server {
	_, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/go-first")
	if err != nil {
		panic(err)
	}

	// Initialize repositories.

	// Initialize domain services.

	return &server{}
}

func (s *server) GetRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middlewares.ResponseAsJson)
	r.Use(middleware.AllowContentType("application/json"))

	r.Get("/", homePage)
	r.Get("/files/{id}", s.fileDetails)
	r.Get("/files", s.fileList)
	r.Post("/files", s.createNewFile)
	r.Delete("/files/{id}", s.deletefile)

	return r
}
