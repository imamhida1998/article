package usecase

import (
	"article/model"
	"article/model/request"
	"article/service/repository"
	"context"
	"fmt"
)

type article struct {
	articleRepo repository.ArticleRepository
}

func NewArticleUsecase(articleRepo repository.ArticleRepository) ArticleUsecase {
	return &article{}
}

func (a *article) CreateArticle(ctx context.Context, params model.Article) error {
	if err := a.articleRepo.CreateArticle(ctx, params); err != nil {
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
