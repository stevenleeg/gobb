package models

import (
	"fmt"
)

type Board struct {
	Id          int64  `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
}

func GetBoard(id int) (error, *Board) {
	db := GetDbSession()
	board := new(Board)
	err := db.SelectOne(board, "SELECT * FROM boards WHERE id=$1", id)

	if err != nil {
		return err, nil
	}

	return nil, board
}

func (board *Board) GetLatestPost() *Post {
	db := GetDbSession()
	latest := &Post{}

	err := db.SelectOne(latest, "SELECT * FROM posts WHERE board_id=$1 AND parent_id IS NULL ORDER BY latest_reply DESC LIMIT 1", board.Id)

	if err != nil {
		fmt.Printf("[error] Could not get latest post in board: (%s)\n", err.Error())
	}

	return latest
}
