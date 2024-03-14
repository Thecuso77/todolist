package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.indexHandler)
	mux.HandleFunc("/api/user", app.controller.GetUser)
	mux.HandleFunc("/api/user/create", app.controller.CreateUser)
	mux.HandleFunc("/api/user/login", app.controller.LoginUser)

	//staticFS, _ := fs.Sub(frontend.StaticFiles, "dist") //Возвращает подкаталоги в dist (объект ) // ui.StaticFiles = embed.FS - функция встраивания файлов в двоичный код программы
	//httpFS := http.FileServer(http.FS(staticFS))        // преобразование embed.Fs в http.FileSystem для проброса в handle
	//mux.Handle("/static/", httpFS)                      // Подкаталоги статических файлов доступны по адресу /static/...
	//
	httpFS := http.FileServer(http.Dir("frontend/dist"))
	mux.Handle("/static/", httpFS)

	return mux
}
