package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"todolist/app/server/db"
	"todolist/app/server/models"
)

type TodoListItemModel struct {
	TodoListItem *db.TodoItemDb
	Users        *db.UsersDb
}

func (tm *TodoListItemModel) LoginUser(w http.ResponseWriter, r *http.Request) {
	var regData models.Users
	// var reqProfile UserRequestProfile

	err := json.NewDecoder(r.Body).Decode(&regData)
	if err != nil {
		panic(err)
	}

	fmt.Println(regData)

	//Откуда брать логин
	err = tm.Users.IsUserExist(regData.Login)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"aliva": true}`)
}

//func (tm *TodoListItemModel) checkItemExist(login string) bool {
//	isExist := tm.TodoListItem.Add(login)
//
//	return isExist
//}
