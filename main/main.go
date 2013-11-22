package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "sirjtaa/controllers"
    "sirjtaa/config"
    "fmt"
    "flag"
)

func main() {
    // Get the config file
    var config_path string
    flag.StringVar(&config_path, "config", "gobb.conf", "Specifies the location of a config file")
    flag.Parse()
    config.GetConfig(config_path)

    // URL Routing!
    r := mux.NewRouter()

    r.HandleFunc("/", controllers.Index)
    r.HandleFunc("/register", controllers.Register)
    r.HandleFunc("/login", controllers.Login)
    r.HandleFunc("/logout", controllers.Logout)
    r.HandleFunc("/admin", controllers.Admin)
    r.PathPrefix("/static/").Handler(http.FileServer(http.Dir("./")))
    http.Handle("/", r)

    fmt.Println("Starting server on port 8080")
    http.ListenAndServe(":8080", nil)
}
