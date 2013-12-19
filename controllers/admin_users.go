package controllers

import (
	"github.com/stevenleeg/gobb/models"
	"github.com/stevenleeg/gobb/utils"
	"github.com/gorilla/mux"
    "database/sql"
	"net/http"
    "strconv"
)

func AdminUsers(w http.ResponseWriter, r *http.Request) {
	current_user := utils.GetCurrentUser(r)
	if current_user == nil || !current_user.IsAdmin() {
		http.NotFound(w, r)
		return
	}

	err := ""
	success := false

	db := models.GetDbSession()
	var users []*models.User
	db.Select(&users, "SELECT * FROM users LIMIT 15")

	utils.RenderTemplate(w, r, "admin_users.html", map[string]interface{}{
		"error":   err,
		"success": success,
		"users":   users,
	}, nil)
}

func AdminUser(w http.ResponseWriter, r *http.Request) {
	current_user := utils.GetCurrentUser(r)
	if current_user == nil || !current_user.IsAdmin() {
		http.NotFound(w, r)
		return
	}

	id_str := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(id_str)

    user, err := models.GetUser(id)

    if err != nil || user == nil {
        http.NotFound(w, r)
        return
    }

    var form_error string
	success := false
	if r.Method == "POST" {
		db := models.GetDbSession()
		user.Avatar = r.FormValue("avatar_url")
		user.UserTitle = r.FormValue("user_title")
		user.StylesheetUrl = sql.NullString{
			Valid:  true,
			String: r.FormValue("stylesheet_url"),
		}
		if r.FormValue("signature") == "" {
			user.Signature = sql.NullString{
				Valid:  false,
				String: r.FormValue("signature"),
			}
		} else {
			user.Signature = sql.NullString{
				Valid:  true,
				String: r.FormValue("signature"),
			}
		}

		// Change hiding settings
		user.HideOnline = false
		if r.FormValue("hide_online") == "1" {
			user.HideOnline = true
		}

		// Update password?
		new_pass := r.FormValue("password_new")
		new_pass2 := r.FormValue("password_new2")
        if len(new_pass) > 0 {
            if len(new_pass) < 5 {
                form_error = "Password must be greater than 4 characters"
            } else if new_pass != new_pass2 {
                form_error = "Passwords didn't match"
            } else {
                user.SetPassword(new_pass)
            }
        }

        group_id, _ := strconv.Atoi(r.FormValue("group_id"))
        user.GroupId = int64(group_id)

		if form_error == "" {
			db.Update(user)
			success = true
		}
	}

	utils.RenderTemplate(w, r, "admin_user.html", map[string]interface{}{
		"error":   form_error,
		"success": success,
		"user":   user,
	}, nil)
}
