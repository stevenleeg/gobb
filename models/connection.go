package models

import (
	"database/sql"
	"fmt"
	"github.com/coopernurse/gorp"
	_ "github.com/lib/pq"
	"github.com/stevenleeg/gobb/config"
)

var db_map *gorp.DbMap

func GetDbSession() *gorp.DbMap {
	if db_map != nil {
		return db_map
	}

	db_username, _ := config.Config.GetString("database", "username")
	db_password, _ := config.Config.GetString("database", "password")
	db_database, _ := config.Config.GetString("database", "database")
	db_hostname, _ := config.Config.GetString("database", "hostname")
	db_port, _ := config.Config.GetString("database", "port")

	if db_port == "" {
		db_port = "5432"
	}

	db, err := sql.Open("postgres",
		"user="+db_username+
			" password="+db_password+
			" dbname="+db_database+
			" host="+db_hostname+
			" port="+db_port+
			" sslmode=disable")

	if err != nil {
		fmt.Printf("Cannot open database! Error: %s\n", err.Error())
		return nil
	}

	db_map = &gorp.DbMap{
		Db:      db,
		Dialect: gorp.PostgresDialect{},
	}

	// TODO: Do we need this every time?
	db_map.AddTableWithName(User{}, "users").SetKeys(true, "Id")
	db_map.AddTableWithName(Board{}, "boards").SetKeys(true, "Id")
	db_map.AddTableWithName(Post{}, "posts").SetKeys(true, "Id")
	db_map.AddTableWithName(View{}, "views").SetKeys(false, "Id")
	db_map.AddTableWithName(Setting{}, "settings").SetKeys(true, "Key")

	return db_map
}
