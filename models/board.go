package models

import (
	"fmt"
	"github.com/stevenleeg/gobb/config"
	"math"
)

type Board struct {
	Id          int64  `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Order   	int    `db:"boardorder"`
}

type BoardLatest struct {
	Op     *Post
	Latest *Post
}

func NewBoard(title, desc string, order int) *Board {
	return &Board{
		Title:       title,
		Description: desc,
		Order:		 order,
	}
}

func UpdateBoard(title, desc string, order int, id int64) *Board {
	return &Board{
		Title:		 title,
		Description: desc,
		Order:		 order,
		Id:			 id,
	}
}

func GetBoard(id int) (*Board, error) {
	db := GetDbSession()
	obj, err := db.Get(&Board{}, id)
	if obj == nil {
		return nil, err
	}

	return obj.(*Board), err
}

func GetBoards() ([]*Board, error) {
	db := GetDbSession()

	var boards []*Board
	_, err := db.Select(&boards, "SELECT * FROM boards ORDER BY boardorder ASC")

	return boards, err
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

	return BoardLatest{
		Op:     op,
		Latest: latest,
	}
}

func (board *Board) GetThreads(page int) ([]*Post, error) {
	db := GetDbSession()
	threads_per_page, err := config.Config.GetInt64("gobb", "threads_per_page")

	var threads []*Post
	i_begin := int64(page) * (threads_per_page - 1)
	_, err = db.Select(&threads, "SELECT * FROM posts WHERE board_id=$1 AND parent_id IS NULL ORDER BY sticky DESC, latest_reply DESC LIMIT $2 OFFSET $3", board.Id, threads_per_page-1, i_begin)

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

// Deletes a board and all of the posts it contains
func (board *Board) Delete() {
	db := GetDbSession()
	db.Exec("DELETE FROM posts WHERE board_id=$1", board.Id)
	db.Delete(board)
}
