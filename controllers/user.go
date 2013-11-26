package controllers

import (
    "github.com/stevenleeg/gobb/utils"
    "github.com/stevenleeg/gobb/models"
	"github.com/gorilla/mux"
    "net/http"
    "strconv"
)

func User(w http.ResponseWriter, r *http.Request) {
    db := models.GetDbSession()
    
	user_id_str := mux.Vars(r)["id"]
    user_id, err := strconv.Atoi(user_id_str)

    if err != nil {
        http.NotFound(w, r)
        return
    }

    user, err := db.Get(&models.User{}, user_id)
    if err != nil {
        http.NotFound(w, r)
        return
    }

    utils.RenderTemplate(w, r, "user.html", map[string]interface{} {
        "user": user,
    }, nil)
}
