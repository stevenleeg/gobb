package controllers

import (
	"github.com/gorilla/mux"
	"github.com/stevenleeg/gobb/models"
	"github.com/stevenleeg/gobb/utils"
	"net/http"
	"strconv"
	"database/sql"
)

func UserSettings(w http.ResponseWriter, r *http.Request) {
	user_id_str := mux.Vars(r)["id"]
	user_id, _ := strconv.Atoi(user_id_str)

	current_user := utils.GetCurrentUser(r)

	if int64(user_id) != current_user.Id {
		http.NotFound(w, r)
		return
	}

	success := false
	if r.Method == "POST" {
		db := models.GetDbSession()
		current_user.Avatar = r.FormValue("avatar_url")
		current_user.StylesheetUrl = sql.NullString{
			Valid: true,
			String: r.FormValue("stylesheet_url"),
		}
		db.Update(current_user)
		success = true
	}

	stylesheet := ""
	if current_user.StylesheetUrl.Valid {
		stylesheet = current_user.StylesheetUrl.String
	}
	
	utils.RenderTemplate(w, r, "user_settings.html", map[string]interface{}{
		"success": success,
		"user_stylesheet": stylesheet,
	}, nil)
}
