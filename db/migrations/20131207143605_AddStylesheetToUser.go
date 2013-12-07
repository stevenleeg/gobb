
package main

import (
	"database/sql"
)

// Up is executed when this migration is applied
func Up_20131207143605(txn *sql.Tx) {
	txn.Exec("ALTER TABLE users ADD COLUMN stylesheet_url varchar")
}

// Down is executed when this migration is rolled back
func Down_20131207143605(txn *sql.Tx) {
	txn.Exec("ALTER TABLE users DROP COLUMN stylesheet_url")
}
