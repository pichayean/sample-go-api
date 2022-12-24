package handlers

import (
	"macus/pkg/models"
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
