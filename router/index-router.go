package router

import (
	"backend_technical_test/database"
	"backend_technical_test/handler"
	"backend_technical_test/middleware"
	"backend_technical_test/repository"
	"backend_technical_test/service"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(r *fiber.App) {
	postRepository := repository.NewPostRepository(database.DB)

	postService := service.NewPostService(postRepository)

	postHandler := handler.NewPostHandler(postService)

	r.Use(middleware.Logger)

	//api
	r.Post("/article", postHandler.CreatePost)
	r.Get("/article/:limit/:offset", postHandler.GetAllPostsByLimitAndOffset)
	r.Get("/article/:id", postHandler.GetPostById)
	r.Put("/article/:id", postHandler.UpdatePost)
	r.Delete("/article/:id", postHandler.DeletePost)
}
