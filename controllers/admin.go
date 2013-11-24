package controllers

import (
	"github.com/stevenleeg/gobb/utils"
	"net/http"
)

func Admin(w http.ResponseWriter, request *http.Request) {
	utils.RenderTemplate(w, request, "admin.html", nil)
}
