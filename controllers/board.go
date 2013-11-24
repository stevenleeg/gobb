package controllers

import (
    "github.com/gorilla/mux"
    "net/http"
    "fmt"
    "sirjtaa/utils"
    "sirjtaa/models"
)

func Board(w http.ResponseWriter, request *http.Request) {
    board := mux.Vars(request)["id"]

    db := models.GetDbSession()
    var threads []models.Post
    _, err := db.Select(&threads, "SELECT * FROM posts WHERE board_id=$1 AND parent_id IS NULL", board)
    if err != nil {
        fmt.Printf("[error] Could not get posts (%s)\n", err.Error())
    }

    utils.RenderTemplate(w, request, "board.html", map[string]interface{} {
        "board": board,
        "threads": threads,
    })
}
