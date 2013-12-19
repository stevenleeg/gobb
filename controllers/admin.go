package controllers

import (
	"github.com/stevenleeg/gobb/models"
	"github.com/stevenleeg/gobb/utils"
	"net/http"
)

func Admin(w http.ResponseWriter, r *http.Request) {
	current_user := utils.GetCurrentUser(r)
	if current_user == nil || !current_user.IsAdmin() {
		http.NotFound(w, r)
		return
	}

	var err error
	success := false
	stylesheet, _ := models.GetStringSetting("theme_stylesheet")
	if r.Method == "POST" {
		stylesheet = r.FormValue("theme_stylesheet")
		models.SetStringSetting("theme_stylesheet", stylesheet)
		success = true
	}

	utils.RenderTemplate(w, r, "admin.html", map[string]interface{}{
		"error":            err,
		"success":          success,
		"theme_stylesheet": stylesheet,
	}, nil)
}
