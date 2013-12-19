package controllers

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stevenleeg/gobb/config"
	"github.com/stevenleeg/gobb/models"
	"github.com/stevenleeg/gobb/utils"
	"net/http"
	"strconv"
	"time"
)

func Thread(w http.ResponseWriter, r *http.Request) {
	page_id_str := r.FormValue("page")
	page_id, err := strconv.Atoi(page_id_str)
	if err != nil {
		page_id = 0
	}

	board_id_str := mux.Vars(r)["board_id"]
	board_id, _ := strconv.Atoi(board_id_str)
	board, err := models.GetBoard(board_id)

	post_id_str := mux.Vars(r)["post_id"]
	post_id, _ := strconv.Atoi(post_id_str)
	err, op, posts := models.GetThread(post_id, page_id)

	var posting_error error

	if r.Method == "POST" {
		db := models.GetDbSession()
		title := r.FormValue("title")
		content := r.FormValue("content")

		current_user := utils.GetCurrentUser(r)
		if current_user == nil {
			http.NotFound(w, r)
			return
		}

        if op.Locked && !current_user.CanModerate() {
			http.NotFound(w, r)
			return
        }

		post := models.NewPost(current_user, board, title, content)
		post.ParentId = sql.NullInt64{int64(post_id), true}
		op.LatestReply = time.Now()

		posting_error = post.Validate()

		if posting_error == nil {
			db.Insert(post)
			db.Update(op)

			if page := post.GetPageInThread(); page != page_id {
				http.Redirect(w, r, fmt.Sprintf("/board/%d/%d?page=%d#post_%d", post.BoardId, op.Id, page, post.Id), http.StatusFound)
			}

			err, op, posts = models.GetThread(post_id, page_id)
		}
	}

	if err != nil {
		http.NotFound(w, r)
		fmt.Printf("[error] Something went wrong in posts (%s)\n", err.Error())
		return
	}

	num_pages := op.GetPagesInThread()

	if page_id > num_pages {
		http.NotFound(w, r)
		return
	}

    var previous_text string
    if posting_error != nil {
        previous_text = r.FormValue("content")
    }

	utils.RenderTemplate(w, r, "thread.html", map[string]interface{}{
		"board":         board,
		"op":            op,
		"posts":         posts,
        "first_page":    (page_id > 0),
		"prev_page":     (page_id > 1),
		"next_page":     (page_id < num_pages - 1),
        "last_page":     (page_id < num_pages),
		"page_id":       page_id,
		"posting_error": posting_error,
        "previous_text": previous_text,
	}, map[string]interface{}{

        "CurrentUserCanModerateThread": func(thread *models.Post) bool {
			current_user := utils.GetCurrentUser(r)
            if current_user == nil {
                return false
            }

            return (current_user.CanModerate() && thread.ParentId.Valid == false)
        },

		"CurrentUserCanDeletePost": func(thread *models.Post) bool {
			current_user := utils.GetCurrentUser(r)
			if current_user == nil {
				return false
			}

			return (current_user.Id == thread.AuthorId) || current_user.CanModerate()
		},

		"CurrentUserCanEditPost": func(post *models.Post) bool {
			current_user := utils.GetCurrentUser(r)
			if current_user == nil {
				return false
			}

			return (current_user.Id == post.AuthorId || current_user.CanModerate())
		},

        "CurrentUserCanModerate": func() bool {
			current_user := utils.GetCurrentUser(r)
			if current_user == nil {
				return false
			}

            return current_user.CanModerate()
        },

		"SignaturesEnabled": func() bool {
			enable_signatures, _ := config.Config.GetBool("gobb", "enable_signatures")
			return enable_signatures
		},

        "ShowReplyBox": func(post *models.Post) bool {
			current_user := utils.GetCurrentUser(r)
            if current_user != nil && (!post.Locked || current_user.CanModerate()) {
                return true
            }
            return false
        },
	})
}
