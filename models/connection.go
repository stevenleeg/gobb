package models

import (
    "github.com/coopernurse/gorp"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "fmt"
)

var db_map *gorp.DbMap

func GetDbSession() *gorp.DbMap {
    if db_map != nil {
        return db_map
    }

    db, err := sql.Open("sqlite3", "/tmp/sirjtaa.db"); 
    if err != nil {
        fmt.Printf("Cannot open database! Error: %s\n", err.Error())
        return nil
    }

    db_map := &gorp.DbMap{
        Db: db,
        Dialect: gorp.SqliteDialect{},
    }

    // TODO: Do we need this every time?
    db_map.AddTableWithName(User{}, "users").SetKeys(true, "Id")
    db_map.CreateTablesIfNotExists()

    return db_map
}
