package models

import (
    "github.com/coopernurse/gorp"
    "database/sql"
    "errors"
    "time"
    "fmt"
)

type Post struct {
    Id        int64         `db:"id"`
    BoardId   int64         `db:"board_id"` 
    ParentId  sql.NullInt64 `db:"parent_id"` 
    Author    *User         `db:"-"`
    AuthorId  int64         `db:"author_id"` 
    Title     string        `db:"title"`
    Content   string        `db:"content"`
    CreatedOn time.Time     `db:"created_on"`
}

func NewPost(author *User, board *Board, title, content string) *Post {
    // TODO: Validation

    post := &Post {
        BoardId: board.Id,
        AuthorId: author.Id,
        Title: title,
        Content: content,
        CreatedOn: time.Now(),
    }

    return post
}

func GetThread(parent_id int) (error, *Post, []*Post) {
    db := GetDbSession()

    op, err := db.Get(Post{}, parent_id)
    if err != nil {
        return errors.New("Parent doesn't exist"), nil, nil
    }

    var child_posts []*Post
    db.Select(&child_posts, "SELECT * FROM posts WHERE parent_id=$1", parent_id)

    return nil, op.(*Post), child_posts
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
