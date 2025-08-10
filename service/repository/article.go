package repository

import (
	"article/model"
	"article/model/request"
	"context"
)

type ArticleRepository interface {
	CreateArticle(ctx context.Context, params model.Article) error
	GetArticle(ctx context.Context, params request.GetListArticle) ([]model.Article, error)
}
