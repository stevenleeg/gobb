package models

import (
    "database/sql"
    "fmt"
    "errors"
    "time"
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

    var child_posts []Post
    var ret_posts []*Post
    
    // Get the initial thread post
    db.Select(&child_posts, "SELECT * FROM posts WHERE parent_id=$1", parent_id)

    for _, post := range child_posts {
        post_ptr := &post
        post.GetAuthor()
        ret_posts = append(ret_posts, post_ptr)
    }

    return nil, op.(*Post), ret_posts
}

func (post *Post) GetAuthor() {
    db := GetDbSession()
    user, _ := db.Get(User{}, post.AuthorId)

    if user == nil {
        fmt.Println("Something went wrong")
    }

    post.Author = user.(*User)
}
