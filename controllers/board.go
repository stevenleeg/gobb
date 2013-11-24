package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stevenleeg/gobb/models"
	"github.com/stevenleeg/gobb/utils"
	"net/http"
	"strconv"
)

func Board(w http.ResponseWriter, r *http.Request) {
	db := models.GetDbSession()

	board_id_str := mux.Vars(r)["id"]
	board_id, _ := strconv.Atoi(board_id_str)
	board, err := db.Get(models.Board{}, board_id)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	var threads []*models.Post
	_, err = db.Select(&threads, "SELECT * FROM posts WHERE board_id=$1 AND parent_id IS NULL ORDER BY latest_reply DESC", board_id)
	if err != nil {
		fmt.Printf("[error] Could not get posts (%s)\n", err.Error())
	}

	utils.RenderTemplate(w, r, "board.html", map[string]interface{}{
		"board":   board,
		"threads": threads,
	})
}
