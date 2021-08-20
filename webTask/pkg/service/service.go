package service

import (
	"webTask"
	"webTask/pkg/repository"
)

type Authorization interface {
	CreateUser(user webTask.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list webTask.TodoList) (int, error)
	GetAll(userId int) ([]webTask.TodoList, error)
	GetById(userId, listId int) (webTask.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input webTask.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, item webTask.TodoItem) (int, error)
	GetAll(userId, listId int) ([]webTask.TodoItem, error)
	GetById(userId, itemId int) (webTask.TodoItem, error)
	Update(userId, itemId int, input webTask.UpdateItemInput) error
	Delete(userId, itemId int) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: newAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
