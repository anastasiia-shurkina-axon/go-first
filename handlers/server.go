package handlers

import (
	"database/sql"

	"github.com/anastasiia-shurkina-axon/go-first/domains/article"
	"github.com/anastasiia-shurkina-axon/go-first/middlewares"
	"github.com/anastasiia-shurkina-axon/go-first/repositories"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Server interface {
	GetRouter() *mux.Router
}

type server struct {
	articleService article.Service
}

func NewServer() Server {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/go-first")
	if err != nil {
		panic(err)
	}

	// Initialize repositories.
	articleRepository := repositories.NewArticleRepository(db)

	// Initialize domain services.
	articleService := article.NewService(articleRepository)

	return &server{
		articleService: articleService,
	}
}

func (s *server) GetRouter() *mux.Router {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage).Methods("GET")

	myRouter.HandleFunc("/articles/{id}", s.articleDetails).Methods("GET")
	myRouter.HandleFunc("/articles", s.createNewArticle).Methods("POST")
	myRouter.HandleFunc("/articles", s.articleList).Methods("GET")

	myRouter.Use(
		middlewares.LogRequest,
		middlewares.ResponseAsJson,
	)
	return myRouter
}
