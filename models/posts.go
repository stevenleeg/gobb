package models

import (
    "database/sql"
    "time"
)

type Post struct {
    Id        int64         `db:"id"`
    BoardId   int64         `db:"board_id"` 
    ParentId  sql.NullInt64 `db:"parent_id"` 
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
