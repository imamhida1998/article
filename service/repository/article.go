package repository

import (
	"article/model"
	"article/model/request"
	"context"
)

type ArticleRepository interface {
	CreateArticle(params model.Article, ctx context.Context) error
	GetArticle(ctx context.Context, params request.GetListArticle) ([]model.Article, error)
}
