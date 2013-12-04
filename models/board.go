package models

import (
    "github.com/stevenleeg/gobb/config"
    "math"
	"fmt"
)

type Board struct {
	Id          int64  `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
}

type BoardLatest struct {
    Op     *Post
    Latest *Post
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

func (board *Board) GetLatestPost() BoardLatest {
	db := GetDbSession()
    op := &Post{}
    latest := &Post{}

	err := db.SelectOne(op, "SELECT * FROM posts WHERE board_id=$1 AND parent_id IS NULL ORDER BY latest_reply DESC LIMIT 1", board.Id)

	if err != nil {
		fmt.Printf("[error] Could not get latest post in board: (%s)\n", err.Error())
	}

	err = db.SelectOne(latest, "SELECT * FROM posts WHERE board_id=$1 AND parent_id=$2 ORDER BY created_on DESC LIMIT 1", board.Id, op.Id)

    if latest.Author == nil {
        latest = nil
    }

    return BoardLatest {
        Op:     op,
        Latest: latest,
    }
}


func (board *Board) GetThreads(page int) ([]*Post, error) {
    db := GetDbSession()
    threads_per_page, err := config.Config.GetInt64("gobb", "threads_per_page")

	var threads []*Post
    i_begin := int64(page) * (threads_per_page - 1)
	_, err = db.Select(&threads, "SELECT * FROM posts WHERE board_id=$1 AND parent_id IS NULL ORDER BY sticky DESC, latest_reply DESC LIMIT $2 OFFSET $3", board.Id, threads_per_page - 1, i_begin)

    return threads, err
}

func (board *Board) GetPagesInBoard() int {
    db := GetDbSession()
    count, err := db.SelectInt("SELECT COUNT(*) FROM posts WHERE board_id=$1 AND parent_id IS NULL", board.Id)

    threads_per_page, err := config.Config.GetInt64("gobb", "threads_per_page")
    
    if err != nil {
        threads_per_page = 30
    }

    return int(math.Floor(float64(count) / float64(threads_per_page)))
}
