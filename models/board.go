package models

type Board struct {
    Id          int64  `db:"id"`
    Title       string `db:"title"`
    Description string `db:"description"`
}
