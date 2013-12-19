package controllers

import (
	"github.com/stevenleeg/gobb/models"
	"github.com/stevenleeg/gobb/utils"
	"net/http"
	"strconv"
)

func AdminBoards(w http.ResponseWriter, r *http.Request) {
	current_user := utils.GetCurrentUser(r)
	if current_user == nil || !current_user.IsAdmin() {
		http.NotFound(w, r)
		return
	}

	db := models.GetDbSession()
	// Creating a board
	if r.Method == "POST" && r.FormValue("create_board") != "" {
		name := r.FormValue("title")
		desc := r.FormValue("description")
		form_order := r.FormValue("order")
		var order int

		if form_order != "" {
			if len(form_order) == 0 {
				order = 1
			} else {
				order, _ = strconv.Atoi(form_order)
			}
		} else {
			order = 1
		}

		board := models.NewBoard(name, desc, order)

		db.Insert(board)
	}

	// Update the boards
	if r.Method == "POST" && r.FormValue("update_boards") != "" {
		err := r.ParseForm()

		// loop through the post data, entries correspond via index in the map
		for i := 0; i < len(r.Form["board_id"]); i++ {
			// basically repeat the process for inserting a board
			form_id, _ := strconv.Atoi(r.Form["board_id"][i])
			id := int64(form_id)
			name := r.Form["name"][i]
			desc := r.Form["description"][i]
			form_order := r.Form["order"][i]
			var order int

			if form_order != "" {
				if len(form_order) == 0 {
					order = 1
				} else {
					order, _ = strconv.Atoi(form_order)
				}
			} else {
				order = 1
			}
			board := models.UpdateBoard(name, desc, order, id)

			db.Update(board)
		}

		if err != nil {
			http.NotFound(w, r)
			return
		}
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
