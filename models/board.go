package models

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
