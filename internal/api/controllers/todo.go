package controllers

import (
	"macus/internal/pkg/models"
	"macus/pkg/services"
	"net/http"

	"github.com/PeteProgrammer/go-automapper"
	"github.com/gin-gonic/gin"
	"gopkg.in/h2non/gentleman.v2"
)

type TodoHandler struct {
	todoService services.TodoService
}

func NewTodoHandler(c *gentleman.Client) *TodoHandler {
	ser := services.NewTodoService(c)
	return &TodoHandler{todoService: ser}
}

// Todos godoc
// @Summary  Get All Todos
// @Schemes
// @Description Get All Todos
// @Tags Todos
// @Accept json
// @Produce json
// @Success 200 {object} models.TodoModel
// @Router /todos [get]
func (h TodoHandler) GetAll(c *gin.Context) {
	todoList, err := h.todoService.GetAll()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	todoModels := make([]models.TodoModel, len(todoList))
	for i := range todoList {
		automapper.Map(todoList[i], &todoModels[i])
	}
	c.JSON(http.StatusOK, todoModels)
}
