package repository

import (
	"article/model"
	"article/model/request"
	"context"
	"database/sql"
	"fmt"
	"time"
)

type articleRepository struct {
	db *sql.DB
}

func NewArticleRepository(db *sql.DB) ArticleRepository {
	return &articleRepository{db}
}

func (repo *articleRepository) CreateArticle(params model.Article, ctx context.Context) error {
	query := `INSERT INTO articles (id,author_id, title, body) 
							VALUES ($1, $2, $3,$4 ) `

	_, err := repo.db.Exec(query, params.Id, params.AuthorId, params.Title, params.Body)
	if err != nil {
		return err
	}

	return nil

}

func (repo *articleRepository) GetArticle(ctx context.Context, params request.GetListArticle) ([]model.Article, error) {
	var args []interface{}
	argIndex := 1

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	sqlBuilder := `
            SELECT id, title, body, author_id, created_at
            FROM articles
            WHERE 1=1
        `
	if params.Query != "" {
		sqlBuilder += fmt.Sprintf(" AND (title ILIKE $%d OR body ILIKE $%d)", argIndex, argIndex+1)
		args = append(args, "%"+params.Query+"%", "%"+params.Query+"%")
		argIndex += 2
	}

	if params.Author != "" {
		sqlBuilder += fmt.Sprintf(" AND author_id ILIKE $%d", argIndex)
		args = append(args, "%"+params.Author+"%")
		argIndex++
	}

	sqlBuilder += " ORDER BY created_at DESC"

	rows, err := repo.db.QueryContext(ctx, sqlBuilder, args...)
	if err != nil {

		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var articles []model.Article
	for rows.Next() {
		var a model.Article
		if err := rows.Scan(&a.Id, &a.Title, &a.Body, &a.AuthorId, &a.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		articles = append(articles, a)
	}

	return articles, nil
}
