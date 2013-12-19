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

	if current_user == nil || int64(user_id) != current_user.Id {
		http.NotFound(w, r)
		return
	}

	success := false
	var form_error string
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

		// Change hiding settings
		current_user.HideOnline = false
		if r.FormValue("hide_online") == "1" {
			current_user.HideOnline = true
		}

		// Update password?
		old_pass := r.FormValue("password_old")
		new_pass := r.FormValue("password_new")
		new_pass2 := r.FormValue("password_new2")
		if old_pass != "" {
			err, user := models.AuthenticateUser(current_user.Username, old_pass)
			if user == nil || err != nil {
				form_error = "Invalid password"
			} else if len(new_pass) < 5 {
                form_error = "Password must be greater than 4 characters"
            } else if new_pass != new_pass2 {
				form_error = "Passwords didn't match"
			} else {
				current_user.SetPassword(new_pass)
                session, _ := utils.GetCookieStore(r).Get(r, "sirsid")
                session.Values["password"] = new_pass
                session.Save(r, w)
			}
		}

		if form_error == "" {
			db.Update(current_user)
			success = true
		}
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
		"error":             form_error,
		"success":           success,
		"user_stylesheet":   stylesheet,
		"user_signature":    signature,
		"enable_signatures": enable_signatures,
	}, nil)
}
