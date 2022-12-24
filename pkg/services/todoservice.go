package services

import (
	"fmt"

	"gopkg.in/h2non/gentleman.v2"
)

type TodoResult struct {
	UserId    int
	Id        int
	Title     string
	Completed bool
}

type TodoService interface {
	GetAll() ([]TodoResult, error)
}

type todoHttpClient struct {
	httpClient *gentleman.Client
}

func NewTodoService(c *gentleman.Client) *todoHttpClient {
	return &todoHttpClient{
		httpClient: c,
	}
}

func (todo *todoHttpClient) GetAll() ([]TodoResult, error) {
	req := todo.httpClient.Request()
	req.Path("/todos")
	res, err := req.Send()
	if err != nil {
		fmt.Printf("Request error: %s\n", err)
		return nil, err
	}
	if !res.Ok {
		fmt.Printf("Invalid server response: %d\n", res.StatusCode)
		return nil, err
	}

	// Reads the whole body and returns it as string
	// fmt.Printf("Body: %s", res.String())
	var resp []TodoResult
	res.JSON(&resp)
	return resp, nil
}
