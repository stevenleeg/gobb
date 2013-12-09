package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stevenleeg/gobb/config"
	"github.com/stevenleeg/gobb/controllers"
	"net/http"
    "go/build"
    "path/filepath"
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
	r.HandleFunc("/admin/boards", controllers.AdminBoards)
	r.HandleFunc("/action/stick", controllers.ActionStickThread)
	r.HandleFunc("/action/delete", controllers.ActionDeleteThread)
	r.HandleFunc("/action/move", controllers.ActionMoveThread)
	r.HandleFunc("/action/edit", controllers.PostEditor)
	r.HandleFunc("/board/{id:[0-9]+}", controllers.Board)
	r.HandleFunc("/board/{board_id:[0-9]+}/new", controllers.PostEditor)
	r.HandleFunc("/board/{board_id:[0-9]+}/{post_id:[0-9]+}", controllers.Thread)
	r.HandleFunc("/user/{id:[0-9]+}", controllers.User)
	r.HandleFunc("/user/{id:[0-9]+}/settings", controllers.UserSettings)

    // Handle static files
    pkg, _ := build.Import("github.com/stevenleeg/gobb/gobb", ".", build.FindOnly)
    static_path := filepath.Join(pkg.SrcRoot, pkg.ImportPath)
	r.PathPrefix("/static/").Handler(http.FileServer(http.Dir(static_path)))

	http.Handle("/", r)

	fmt.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", nil)
}
