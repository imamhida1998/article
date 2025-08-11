package usecase

import (
	"article/model"
	"article/model/request"
	"article/service/repository"
	"context"
	"fmt"
	"github.com/google/uuid"
)

type article struct {
	articleRepo repository.ArticleRepository
}

func NewArticleUsecase(articleRepo repository.ArticleRepository) ArticleUsecase {
	return &article{
		articleRepo: articleRepo,
	}
}

func (a *article) CreateArticle(params model.Article, ctx context.Context) error {
	id, _ := uuid.NewV7()
	params.Id = id.String()
	err := a.articleRepo.CreateArticle(params, ctx)
	if err != nil {
		return err
	}
	return nil

}

func (a *article) GetArticle(ctx context.Context, params request.GetListArticle) ([]model.Article, error) {
	article, err := a.articleRepo.GetArticle(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to get article: %w", err)
	}

	if article == nil {
		return nil, fmt.Errorf("article not found")
	}

	return article, nil

}
