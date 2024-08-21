package handler

import (
	"backend_technical_test/entities/request"
	"backend_technical_test/entities/response"
	"backend_technical_test/service"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type postHandler struct {
	postService service.PostService
}

func NewPostHandler(postService service.PostService) *postHandler {
	return &postHandler{postService}
}

func (h *postHandler) CreatePost(ctx *fiber.Ctx) error {
	request := new(request.PostRequest)
	log.Println(request)

	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	// VALIDATE REQUEST

	validate := validator.New()
	errValidate := validate.Struct(request)

	if errValidate != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": errValidate.Error(),
		})
	}
	post, err := h.postService.CreatePost(*request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	return ctx.Status(fiber.StatusCreated).JSON(response.WebResponse{
		Status:  "success",
		Message: "Post created successfully",
		Data:    post,
	})
}

func (h *postHandler) GetAllPostsByLimitAndOffset(ctx *fiber.Ctx) error {
	limit, _ := ctx.ParamsInt("limit")
	offset, _ := ctx.ParamsInt("offset")

	// fmt.Println(limit, offset)

	posts, err := h.postService.GetAllPostsByLimitAndOffset(limit, offset)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.WebResponseWithLimitAndOffset{
		Status:  "success",
		Message: "Posts fetched successfully",
		Data:    posts,
		Limit:   limit,
		Offset:  offset,
	})

}

func (h *postHandler) GetPostById(ctx *fiber.Ctx) error {

	id, _ := ctx.ParamsInt("id")
	post, err := h.postService.GetPostById(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": post})
}

func (h *postHandler) UpdatePost(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")

	request := new(request.PostRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	validate := validator.New()
	errValidate := validate.Struct(request)

	if errValidate != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": errValidate.Error(),
		})
	}
	post, err := h.postService.UpdatePost(id, *request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.WebResponse{
		Status:  "success",
		Message: "Post updated successfully",
		Data:    post,
	})

}

func (h *postHandler) DeletePost(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")
	_, err := h.postService.DeletePost(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.WebResponse{
		Status:  "success",
		Message: "Post deleted successfully",
		Data:    "null",
	})
}
