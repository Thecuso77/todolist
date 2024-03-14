package controllers

import (
	"todolist/app/server/models"
)

type TodoListItemModel struct {
	TodoListItem *models.TodoItemDb
}

func (tm *TodoListItemModel) checkItemExist(login string) bool {
	isExist := tm.TodoListItem.Add(login)

	return isExist
}
