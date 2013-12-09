package controllers

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/stevenleeg/gobb/models"
	"github.com/stevenleeg/gobb/utils"
    "fmt"
	"net/http"
	"strconv"
	"time"
    "github.com/stevenleeg/gobb/config"
)

func Thread(w http.ResponseWriter, r *http.Request) {
    enable_signatures, _ := config.Config.GetBool("gobb", "enable_signatures")

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

	if r.Method == "POST" {
		db := models.GetDbSession()
		title := r.FormValue("title")
		content := r.FormValue("content")

		current_user := utils.GetCurrentUser(r)
		if current_user == nil {
			http.NotFound(w, r)
			return
		}

		post := models.NewPost(current_user, board, title, content)
		post.ParentId = sql.NullInt64{int64(post_id), true}
		op.LatestReply = time.Now()
		db.Insert(post)
		db.Update(op)

        if page := post.GetPageInThread(); page != page_id {
            http.Redirect(w, r, fmt.Sprintf("/board/%d/%d?page=%d#post_%d", post.BoardId, op.Id, page, post.Id), http.StatusFound)
        }

		err, op, posts = models.GetThread(post_id, page_id)
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

	utils.RenderTemplate(w, r, "thread.html", map[string]interface{}{
		"board": board,
		"op":    op,
		"posts": posts,
        "prev_page": (page_id != 0),
        "next_page": (page_id < num_pages),
        "page_id": page_id,
        "enable_signatures": enable_signatures,
	}, map[string]interface{}{

        "CurrentUserCanDeletePost": func(thread *models.Post) bool {
            current_user := utils.GetCurrentUser(r)
            if current_user == nil {
                return false
            }
            
            return (current_user.Id == thread.AuthorId && !thread.ParentId.Valid) || current_user.CanModerate()
        },

        "CurrentUserCanStickyThread": func(thread *models.Post) bool {
            current_user := utils.GetCurrentUser(r)
            if current_user == nil {
                return false
            }

            return (current_user.CanModerate() && !thread.ParentId.Valid)
        },

        "CurrentUserCanMoveThread": func(thread *models.Post) bool {
            current_user := utils.GetCurrentUser(r)
            if current_user == nil {
                return false
            }

            return (current_user.CanModerate() && !thread.ParentId.Valid)
        },

        "CurrentUserCanEditPost": func(post *models.Post) bool {
            current_user := utils.GetCurrentUser(r)
            if current_user == nil {
                return false
            }

            return (current_user.Id == post.AuthorId || current_user.CanModerate())
        },
    })
}
