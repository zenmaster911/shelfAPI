package service

import (
	todo "github.com/zenmaster911/shelfAPI"
	"github.com/zenmaster911/shelfAPI/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodolistService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list todo.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}
