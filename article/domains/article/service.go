package article

import (
	"github.com/anastasiia-shurkina-axon/go-first/article/models"
	"github.com/anastasiia-shurkina-axon/go-first/article/repositories"
)

type Service interface {
	Create(a *models.Article) (*models.Article, error)
	Read(id int) (*models.Article, error)
	List() ([]*models.Article, error)
	Delete(id int) error
}

type service struct {
	repo repositories.ArticleRepository
}

func NewService(repo repositories.ArticleRepository) Service {
	return &service{repo}
}

func (s *service) Create(a *models.Article) (*models.Article, error) {
	article, err := s.repo.Create(a.Title, a.Desc, a.Content)

	return article, err
}

func (s *service) Read(id int) (*models.Article, error) {
	article, err := s.repo.Read(id)

	return article, err
}

func (s *service) List() ([]*models.Article, error) {
	articles, err := s.repo.List()

	return articles, err
}

func (s *service) Delete(id int) error {
	err := s.repo.Delete(id)

	return err
}
