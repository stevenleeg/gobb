package models

import (
	"fmt"
	"github.com/lib/pq"
	"github.com/stevenleeg/gobb/config"
	"math"
	"time"
)

type Board struct {
	Id          int64  `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Order       int    `db:"ordering"`
}

type BoardLatest struct {
	Op     *Post
	Latest *Post
}

type JoinBoardView struct {
    Board       *Board      `db:"-"`
	Id          int64       `db:"id"`
	Title       string      `db:"title"`
	Description string      `db:"description"`
	Order       int         `db:"ordering"`
	ViewedOn    pq.NullTime `db:"viewed_on"`
}

type JoinThreadView struct {
	Thread      *Post       `db:"-"`
	Id          int64       `db:"id"`
	BoardId     int64       `db:"board_id"`
	Author      *User       `db:"-"`
	AuthorId    int64       `db:"author_id"`
	Title       string      `db:"title"`
	CreatedOn   time.Time   `db:"created_on"`
	LatestReply time.Time   `db:"latest_reply"`
	Sticky      bool        `db:"sticky"`
	Locked      bool        `db:"locked"`
	ViewedOn    pq.NullTime `db:"viewed_on"`
}

func NewBoard(title, desc string, order int) *Board {
	return &Board{
		Title:       title,
		Description: desc,
		Order:       order,
	}
}

func UpdateBoard(title, desc string, order int, id int64) *Board {
	return &Board{
		Title:       title,
		Description: desc,
		Order:       order,
		Id:          id,
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
	_, err := db.Select(&boards, "SELECT * FROM boards ORDER BY ordering ASC")

	return boards, err
}

func GetBoardsUnread(user *User) ([]*JoinBoardView, error) {
	db := GetDbSession()

    user_id := int64(-1)
    if user != nil {
        user_id = user.Id
    }

	var boards []*JoinBoardView
	_, err := db.Select(&boards, `
        SELECT
            boards.*,
            views.time AS viewed_on
        FROM boards
        LEFT OUTER JOIN views ON
            views.post_id=(SELECT id FROM posts WHERE board_id=boards.id AND parent_id IS NULL ORDER BY latest_reply DESC LIMIT 1) AND
            views.user_id=$1
        ORDER BY
            ordering ASC
    `, user_id)

    for i := range boards {
        if user_id == -1 {
            boards[i].ViewedOn = pq.NullTime{Time: time.Now(), Valid: true}
        }

        boards[i].Board = &Board{
            Id: boards[i].Id,
        }
    }
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

func (board *Board) GetThreads(page int, user *User) ([]*JoinThreadView, error) {
	db := GetDbSession()
	threads_per_page, err := config.Config.GetInt64("gobb", "threads_per_page")

	var threads []*JoinThreadView
	i_begin := int64(page) * (threads_per_page - 1)

	user_id := int64(-1)
	if user != nil {
		user_id = user.Id
	}
	_, err = db.Select(&threads, `
        SELECT 
            posts.id, 
            posts.author_id,
            posts.title,
            posts.created_on,
            posts.latest_reply,
            posts.sticky,
            posts.locked,
            posts.board_id,
            views.time AS viewed_on 
        FROM posts
        LEFT OUTER JOIN views ON 
            posts.id=views.post_id AND
            views.user_id=$4
        WHERE
            board_id=$1 AND
            parent_id IS NULL
        ORDER BY
            sticky DESC,
            latest_reply DESC
        LIMIT $2 OFFSET $3
    `, board.Id, threads_per_page-1, i_begin, user_id)

	for i := range threads {
        if user_id == -1 {
            threads[i].ViewedOn = pq.NullTime{Time: time.Now(), Valid: true}
        }

		obj, _ := db.Get(&User{}, threads[i].AuthorId)
		user := obj.(*User)
		threads[i].Author = user
		threads[i].Thread = &Post{
			Id: threads[i].Id,
		}
	}

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
