
package main

import (
	"database/sql"
)

// Up is executed when this migration is applied
func Up_20131209104951(txn *sql.Tx) {
    txn.Exec("ALTER TABLE users ADD COLUMN signature varchar")
}

// Down is executed when this migration is rolled back
func Down_20131209104951(txn *sql.Tx) {
    txn.Exec("ALTER TABLE users DROP COLUMN signature")
}
