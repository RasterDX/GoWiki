package main

import (
	"../handlers"
	pghandler "../handlers/pages"
	"github.com/jetflight/src/db"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", pghandler.Handler)
	http.Handle("/res/", http.StripPrefix("/res/", http.FileServer(http.Dir("res"))))
	http.HandleFunc("/view/", handlers.MakeHandler(pghandler.ViewHandler))
	http.HandleFunc("/edit/", handlers.MakeHandler(pghandler.EditHandler))
	http.HandleFunc("/save/", handlers.MakeHandler(pghandler.SaveHandler))
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func setupDb() {
	db.CreateArticles()
}