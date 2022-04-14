package repositories

import (
	"database/sql"

	"github.com/anastasiia-shurkina-axon/go-first/models"
)

type ArticleRepository interface {
	List() ([]*models.Article, error)
	Read(id int) (*models.Article, error)
	Create(title string, desc string, content string) (*models.Article, error)
	Delete(id int) error
	Update(article *models.Article, title string, desc string, content string) (*models.Article, error)
}

type articleRepository struct {
	db *sql.DB
}

func NewArticleRepository(db *sql.DB) ArticleRepository {
	return &articleRepository{db}
}

func (r *articleRepository) List() ([]*models.Article, error) {
	rows, err := r.db.Query("select id, title, description, content from articles")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var articles []*models.Article
	for rows.Next() {
		art := &models.Article{}
		err := rows.Scan(&art.Id, &art.Title, &art.Desc, &art.Content)
		if err != nil {
			panic(err)
		}
		articles = append(articles, art)
	}
	return articles, nil
}

func (r *articleRepository) Read(id int) (*models.Article, error) {
	rows, err := r.db.Query("select id, title, description, content from articles where id = ?", id)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	article := &models.Article{}
	for rows.Next() {
		err := rows.Scan(&article.Id, &article.Title, &article.Desc, &article.Content)
		if err != nil {
			panic(err)
		}
	}
	return article, nil
}

func (r *articleRepository) Create(title string, desc string, content string) (*models.Article, error) {
	stmt, err := r.db.Prepare("insert into articles(title, description, content) values (?, ?, ?)")
	if err != nil {
		return nil, err
	}
	res, err := stmt.Exec(title, desc, content)
	if err != nil {
		return nil, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	article := models.Article{
		Id:      int(lastId),
		Title:   title,
		Desc:    desc,
		Content: content,
	}
	return &article, err
}

func (r *articleRepository) Delete(id int) error {
	return nil
}

func (r *articleRepository) Update(article *models.Article, title string, desc string, content string) (*models.Article, error) {
	stmt, err := r.db.Prepare("update articles set title = ?, description = ?, content = ? where id = ?")
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(title, desc, content, article.Id)
	if err != nil {
		return nil, err
	}

	article.Title = title
	article.Desc = desc
	article.Content = content

	return article, err
}
