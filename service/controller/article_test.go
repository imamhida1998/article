package controller_test

import (
	"article/model"
	"article/model/request"
	"article/service/controller"
	"bytes"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockArticleUsecase struct{}

func (m *mockArticleUsecase) CreateArticle(params model.Article, ctx context.Context) error {
	return nil
}

func (m *mockArticleUsecase) GetArticle(ctx context.Context, params request.GetListArticle) ([]model.Article, error) {
	return []model.Article{
		{Id: "1", AuthorId: "author1", Title: "Test Article", Body: "Body"},
	}, nil
}

func TestCreateArticle_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	ctrl := controller.NewArticleController(&mockArticleUsecase{})
	r.POST("/articles", ctrl.CreateArticle)

	article := model.Article{Id: "1", AuthorId: "author1", Title: "Test", Body: "Body"}
	body, _ := json.Marshal(article)
	req, _ := http.NewRequest("POST", "/articles", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestGetArticle_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	ctrl := controller.NewArticleController(&mockArticleUsecase{})
	r.GET("/articles", ctrl.GetArticle)

	article := model.Article{Id: "1", AuthorId: "author1", Title: "Test", Body: "Body"}
	body, _ := json.Marshal(article)
	req, _ := http.NewRequest("GET", "/articles", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}
