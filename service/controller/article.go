package controller

import (
	"article/model"
	"article/model/request"
	"article/service/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type article struct {
	Usecase usecase.ArticleUsecase
}

func NewArticleController(usecase usecase.ArticleUsecase) article {
	return article{
		Usecase: usecase,
	}
}

func (article *article) CreateArticle(ctx *gin.Context) {
	var params model.Article
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	if err := article.Usecase.CreateArticle(params, nil); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, "Successfully created article")
	return

}

func (article *article) GetArticle(ctx *gin.Context) {
	var params request.GetListArticle

	params.Query = ctx.Query("query")
	params.Author = ctx.Query("author")
	if err := ctx.ShouldBindQuery(&params); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid query parameters"})
		return
	}
	articles, err := article.Usecase.GetArticle(ctx, params)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if articles == nil {
		ctx.JSON(404, gin.H{"error": "No articles found"})
		return
	}
	ctx.JSON(http.StatusOK, articles)
	return

}
