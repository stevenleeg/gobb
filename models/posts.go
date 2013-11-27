package models

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/coopernurse/gorp"
	"time"
    "github.com/stevenleeg/gobb/config"
)

type Post struct {
	Id          int64         `db:"id"`
	BoardId     int64         `db:"board_id"`
	ParentId    sql.NullInt64 `db:"parent_id"`
	Author      *User         `db:"-"`
	AuthorId    int64         `db:"author_id"`
	Title       string        `db:"title"`
	Content     string        `db:"content"`
	CreatedOn   time.Time     `db:"created_on"`
	LatestReply time.Time     `db:"latest_reply"`
    Sticky      bool          `db:"sticky"`
}

func NewPost(author *User, board *Board, title, content string) *Post {
	// TODO: Validation

	post := &Post{
		BoardId:   board.Id,
		AuthorId:  author.Id,
		Title:     title,
		Content:   content,
		CreatedOn: time.Now(),
        Sticky:    false,
	}

	return post
}

func GetThread(parent_id, page_id int) (error, *Post, []*Post) {
	db := GetDbSession()

	op, err := db.Get(Post{}, parent_id)
	if err != nil || op == nil {
		return errors.New("[error] Could not get parent (" + err.Error() + ")"), nil, nil
	}

    posts_per_page, err := config.Config.GetInt64("gobb", "posts_per_page")
    if err != nil {
        posts_per_page = 15
    }

    i_begin := int64(page_id) * (posts_per_page - 1)

	var child_posts []*Post
	db.Select(&child_posts, "SELECT * FROM posts WHERE parent_id=$1 ORDER BY created_on ASC LIMIT $2 OFFSET $3", parent_id, posts_per_page - 1, i_begin)

	return nil, op.(*Post), child_posts
}

func GetPostCount() (int64, error) {
    db := GetDbSession()

    count, err := db.SelectInt("SELECT COUNT(*) FROM posts")
    if err != nil {
        fmt.Printf("[error] Error selecting post count (%s)\n", err.Error())
        return 0, errors.New("Database error: " + err.Error())
    }

    return count, nil
}

func (post *Post) PostGet(s gorp.SqlExecutor) error {
	db := GetDbSession()
	user, _ := db.Get(User{}, post.AuthorId)

	if user == nil {
		return errors.New("Could not find post's author")
	}

	post.Author = user.(*User)

	return nil
}

// This is used primarily for threads. It will find the latest
// post in a thread, allowing for things like "last post was 10
// minutes ago.
func (post *Post) GetLatestPost() *Post {
	db := GetDbSession()
	latest := &Post{}

	err := db.SelectOne(latest, "SELECT * FROM posts WHERE parent_id=$1 ORDER BY created_on DESC LIMIT 1", post.Id)

	if err != nil {
		fmt.Printf("[error] Could not get latest post: (%s)\n", err.Error())
	}

	return latest
}

func (post *Post) GetPagesInThread() int {
    db := GetDbSession()
    count, err := db.SelectInt("SELECT COUNT(*) FROM posts WHERE parent_id=$1", post.Id)

    posts_per_page, err := config.Config.GetInt64("gobb", "posts_per_page")
    
    if err != nil {
        posts_per_page = 15
    }

    return int((count - 1) / posts_per_page)
}

func (post *Post) DeleteAllChildren() error {
	db := GetDbSession()

    _, err := db.Exec("DELETE FROM posts WHERE parent_id=$1", post.Id)
    return err
}
