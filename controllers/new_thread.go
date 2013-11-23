package controllers

import (
    "github.com/gorilla/mux"
    "net/http"
    "strconv"
    "fmt"
    "sirjtaa/utils"
    "sirjtaa/models"
)

func NewThread(w http.ResponseWriter, r *http.Request) {
    db := models.GetDbSession()
    board_id_str := mux.Vars(r)["id"]
    board_id, _ := strconv.Atoi(board_id_str)
    err, board := models.GetBoard(board_id)

    if err != nil {
        http.NotFound(w, r)
        return
    }

    current_user := utils.GetCurrentUser(r)
    if current_user == nil {
        http.NotFound(w, r)
        return
    }

    if r.Method == "POST" {
        title := r.FormValue("title")
        content := r.FormValue("content")

        post := models.NewPost(current_user, board, title, content)
        db.Insert(post)

        http.Redirect(w, r, fmt.Sprintf("/board/%d/%d", board.Id, post.Id), http.StatusFound)
        return
    }

    utils.RenderTemplate(w, r, "new_thread.html", map[string]interface{} {
        "board": board,
    })
}
