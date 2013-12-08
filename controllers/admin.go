package controllers

import (
	"github.com/stevenleeg/gobb/utils"
	"github.com/stevenleeg/gobb/models"
	"net/http"
)

func Admin(w http.ResponseWriter, r *http.Request) {
    current_user := utils.GetCurrentUser(r)
    if !current_user.IsAdmin() {
        http.NotFound(w, r)
        return
    }

	utils.RenderTemplate(w, r, "admin.html", nil, nil)
}

func AdminBoards(w http.ResponseWriter, r *http.Request) {
    current_user := utils.GetCurrentUser(r)
    if !current_user.IsAdmin() {
        http.NotFound(w, r)
        return
    }

    db := models.GetDbSession()
    // Creating a board
    if r.Method == "POST" && r.FormValue("create_board") != "" {
        name := r.FormValue("title")
        desc := r.FormValue("description")

        board := models.NewBoard(name, desc)

        db.Insert(board)
    }

    // Delete a board
    if id := r.FormValue("delete"); id != "" {
        obj, _ := db.Get(&models.Board{}, id)

        if obj == nil {
			http.NotFound(w, r)
            return
        }

        board := obj.(*models.Board)
        board.Delete()
    }

    boards, _ := models.GetBoards()

	utils.RenderTemplate(w, r, "admin_boards.html", map[string]interface{}{
        "boards": boards,
    }, nil)
}
