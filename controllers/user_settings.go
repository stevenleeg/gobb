package controllers

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/stevenleeg/gobb/config"
	"github.com/stevenleeg/gobb/models"
	"github.com/stevenleeg/gobb/utils"
	"net/http"
	"strconv"
)

func UserSettings(w http.ResponseWriter, r *http.Request) {
	enable_signatures, _ := config.Config.GetBool("gobb", "enable_signatures")

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
		current_user.UserTitle = r.FormValue("user_title")
		current_user.StylesheetUrl = sql.NullString{
			Valid:  true,
			String: r.FormValue("stylesheet_url"),
		}
		if r.FormValue("signature") == "" {
			current_user.Signature = sql.NullString{
				Valid:  false,
				String: r.FormValue("signature"),
			}
		} else {
			current_user.Signature = sql.NullString{
				Valid:  true,
				String: r.FormValue("signature"),
			}
		}

        current_user.HideOnline = false
        if r.FormValue("hide_online") == "1" {
            current_user.HideOnline = true
        }

		db.Update(current_user)
		success = true
	}

	stylesheet := ""
	if current_user.StylesheetUrl.Valid {
		stylesheet = current_user.StylesheetUrl.String
	}
	signature := ""
	if current_user.Signature.Valid {
		signature = current_user.Signature.String
	}

	utils.RenderTemplate(w, r, "user_settings.html", map[string]interface{}{
		"success":           success,
		"user_stylesheet":   stylesheet,
		"user_signature":    signature,
		"enable_signatures": enable_signatures,
	}, nil)
}
