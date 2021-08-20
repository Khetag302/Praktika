package repository

import (
	"github.com/jmoiron/sqlx"
	"webTask"
)

type Authorization interface {
	CreateUser(user webTask.User) (int, error)
	GetUser(username, password string) (webTask.User, error)
}

type TodoList interface {
	Create(userId int, list webTask.TodoList) (int, error)
	GetAll(userId int) ([]webTask.TodoList, error)
	GetById(userId, listId int) (webTask.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input webTask.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item webTask.TodoItem) (int, error)
	GetAll(userId, listId int) ([]webTask.TodoItem, error)
	GetById(userId, itemId int) (webTask.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input webTask.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
