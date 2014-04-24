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
	favicon, _ := models.GetStringSetting("favicon_url")
	current_template, _ := models.GetStringSetting("template")

	if r.Method == "POST" {
		stylesheet = r.FormValue("theme_stylesheet")
		favicon = r.FormValue("favicon_url")
		current_template = r.FormValue("template")
		models.SetStringSetting("theme_stylesheet", stylesheet)
		models.SetStringSetting("favicon_url", favicon)
		models.SetStringSetting("template", current_template)
		success = true
	}

	utils.RenderTemplate(w, r, "admin.html", map[string]interface{}{
		"error":            err,
		"success":          success,
		"theme_stylesheet": stylesheet,
		"favicon_url":      favicon,
		"current_template": current_template,
		"templates":        utils.ListTemplates(),
	}, map[string]interface{}{
		"IsCurrentTemplate": func(name string) bool {
			return name == current_template
		},
	})
}
