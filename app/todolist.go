package main

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
	"todolist/app/server"
	"todolist/app/server/controllers"
	"todolist/app/server/db"
)

type application struct {
	controller *controllers.TodoListItemModel
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	log.Info("API Health is OK")
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"aliva": true}`)
}

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
}

func main() {
	log.Info("Starting Todolist API")
	router := mux.NewRouter()
	router.HandleFunc("/healthz", Healthz).Methods("GET")

	server.InitMySqlDb()

	/*
	 Если добавить авторизацию и модель пользователя, как она может передаваться к
	*/
	app := &application{
		controller: &controllers.TodoListItemModel{
			TodoListItem: &db.TodoItemDb{MySqlDB: server.MySqlDB},
			Users:        &db.UsersDb{MySqlDB: server.MySqlDB},
		},
	}

	srv := &http.Server{
		Addr:        ":8080",
		Handler:     app.routes(),
		IdleTimeout: time.Minute,
	}

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
