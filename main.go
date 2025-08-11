package main

import (
	"article/config"
	"article/lib"
	"article/service/controller"
	"article/service/repository"
	"article/service/usecase"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	route := gin.Default()
	set := config.Config{}
	set.CatchError(set.InitEnv())
	Database := set.GetDBConfig()
	db, err := lib.ConnectiontoPostgreSQL(Database)
	if err != nil {
		log.Println(err)
		return
	}

	ArticleRepo := repository.NewArticleRepository(db)
	ArtiUsecase := usecase.NewArticleUsecase(ArticleRepo)
	ArticleController := controller.NewArticleController(ArtiUsecase)
	route.Use(gin.Recovery())
	route.POST("/article", ArticleController.CreateArticle)
	route.GET("/article", ArticleController.GetArticle)

	if err := route.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	} else {
		log.Println("Server is running on port 8080")
	}
}
