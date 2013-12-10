package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stevenleeg/gobb/config"
	"github.com/stevenleeg/gobb/controllers"
	"github.com/stevenleeg/gobb/utils"
	"go/build"
	"net/http"
	"path/filepath"
)

func main() {
	// Get the config file
	var config_path string
	flag.StringVar(&config_path, "config", "gobb.conf", "Specifies the location of a config file")
	run_migrations := flag.Bool("migrate", false, "Runs database migrations")
	ign_migrations := flag.Bool("ignore-migrations", false, "Ignores an out of date database and runs the server anyways")
	flag.Parse()
	config.GetConfig(config_path)

	pkg, _ := build.Import("github.com/stevenleeg/gobb/gobb", ".", build.FindOnly)

	// Do we need to run migrations?
	latest_db_version, migrations, err := utils.GetMigrationInfo()
	if len(migrations) != 0 && *run_migrations {
		fmt.Println("[notice] Running database migrations:\n")
		err = utils.RunMigrations(latest_db_version)
		if err != nil {
			fmt.Printf("[error] Could not run migrations (%s)\n", err.Error())
			return
		}

		fmt.Println("\n[notice] Database migration successful!")
	} else if len(migrations) != 0 && !(*ign_migrations) {
		fmt.Println("Your database appears to be out of date. Please run migrations with --migrate or ignore this message with --ignore-migrations")
		return
	}

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
	static_path := filepath.Join(pkg.SrcRoot, pkg.ImportPath)
	r.PathPrefix("/static/").Handler(http.FileServer(http.Dir(static_path)))

	// User provided static files
	static_path, err = config.Config.GetString("gobb", "base_path")
	if err == nil {
		r.PathPrefix("/assets/").Handler(http.FileServer(http.Dir(static_path)))
	}

	http.Handle("/", r)

	fmt.Println("[notice] Starting server on port 8080")
	http.ListenAndServe(":8080", nil)
}
