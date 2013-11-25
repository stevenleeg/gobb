package controllers

import (
	"github.com/stevenleeg/gobb/utils"
    "github.com/stevenleeg/gobb/models"
    "strconv"
    "fmt"
	"net/http"
)

func ActionStickThread(w http.ResponseWriter, r *http.Request) {
    user := utils.GetCurrentUser(r)
    if !user.CanModerate() {
        http.NotFound(w, r)
        return
    }

    thread_id_str := r.FormValue("thread_id")
    thread_id, err := strconv.Atoi(thread_id_str)

    if err != nil {
        http.NotFound(w, r)
        return 
    }

    db := models.GetDbSession()
    obj, err := db.Get(&models.Post{}, thread_id)
    thread := obj.(*models.Post)

    if thread == nil || err != nil {
        http.NotFound(w, r)
        return
    }

    thread.Sticky = !(thread.Sticky)
    db.Update(thread)

    http.Redirect(w, r, fmt.Sprintf("/board/%d/%d", thread.BoardId, thread.Id), http.StatusFound)
}

func ActionDeleteThread(w http.ResponseWriter, r *http.Request) {
    user := utils.GetCurrentUser(r)
    thread_id_str := r.FormValue("thread_id")
    thread_id, err := strconv.Atoi(thread_id_str)

    if err != nil {
        http.NotFound(w, r)
        return 
    }

    db := models.GetDbSession()
    obj, err := db.Get(&models.Post{}, thread_id)
    thread := obj.(*models.Post)

    if thread == nil || err != nil {
        http.NotFound(w, r)
        return
    }

    if (thread.AuthorId != user.Id && !user.CanModerate()) || thread.ParentId.Valid {
        http.NotFound(w, r)
        return
    }

    thread.DeleteAllChildren()
    db.Delete(thread)

    http.Redirect(w, r, fmt.Sprintf("/board/%d", thread.BoardId), http.StatusFound)
}
