package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "sirjtaa/controllers"
    "fmt"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/", controllers.Index)
    r.HandleFunc("/register", controllers.Register)
    r.HandleFunc("/login", controllers.Login)
    r.HandleFunc("/logout", controllers.Logout)
    r.PathPrefix("/static/").Handler(http.FileServer(http.Dir("./")))
    http.Handle("/", r)

    fmt.Println("Starting server on port 8080")
    http.ListenAndServe(":8080", nil)
}
