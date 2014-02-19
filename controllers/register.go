package controllers

import (
	"fmt"
	"github.com/stevenleeg/gobb/models"
	"github.com/stevenleeg/gobb/utils"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if utils.GetCurrentUser(r) != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")
		confirm := r.FormValue("password2")

		var error string
		if password != confirm {
			error = "Passwords don't match"
		}

		// See if a user with this name already exists
		db := models.GetDbSession()
		count, err := db.SelectInt("SELECT COUNT(*) FROM users WHERE username=$1", username)
		if count > 0 || err != nil {
			error = "This username is already taken."
		}

        if len(username) < 3 {
            error = "Username must be greater than 3 characters."
        }

		if error != "" {
			utils.RenderTemplate(w, r, "register.html", map[string]interface{}{
				"error": error,
			}, nil)
			return
		}

		// We're good, let's make it
		user := models.NewUser(username, password)
		err = db.Insert(user)

		if err != nil {
			fmt.Printf("[error] Could not insert user (%s)\n", err.Error())
			return;
		}

		// Adminify the first user
		id, err := db.SelectInt("SELECT lastval()")
		if err == nil && id == 1 {
			user.GroupId = 2
			count, err = db.Update(user)

			if err != nil {
				fmt.Printf("[error] Could not adminify user (%s)\n", err.Error())
				return;
			}
		}

		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	utils.RenderTemplate(w, r, "register.html", nil, nil)
}
