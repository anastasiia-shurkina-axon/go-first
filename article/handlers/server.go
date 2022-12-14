package handlers

import (
	"database/sql"

	"github.com/anastasiia-shurkina-axon/go-first/article/domains/article"
	"github.com/anastasiia-shurkina-axon/go-first/article/middlewares"
	"github.com/anastasiia-shurkina-axon/go-first/article/repositories"
	_ "github.com/go-sql-driver/mysql"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server interface {
	GetRouter() *chi.Mux
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

func (s *server) GetRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middlewares.ResponseAsJson)
	r.Use(middleware.AllowContentType("application/json"))

	r.Get("/", homePage)
	r.Get("/articles/{id}", s.articleDetails)
	r.Get("/articles", s.articleList)
	r.Post("/articles", s.createNewArticle)
	r.Delete("/articles/{id}", s.deleteArticle)

	return r
}
