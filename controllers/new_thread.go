package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stevenleeg/gobb/models"
	"github.com/stevenleeg/gobb/utils"
	"net/http"
	"strconv"
	"time"
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
		post.LatestReply = time.Now()
		err := db.Insert(post)

        if err != nil {
            fmt.Printf("[error] Could not create new thread (%s)", err.Error())
            return
        }

		http.Redirect(w, r, fmt.Sprintf("/board/%d/%d", board.Id, post.Id), http.StatusFound)
		return
	}

	utils.RenderTemplate(w, r, "new_thread.html", map[string]interface{}{
		"board": board,
	})
}
