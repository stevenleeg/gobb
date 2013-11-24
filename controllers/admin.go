package controllers

import (
    "net/http"
    "github.com/stevenleeg/gobb/utils"
)

func Admin(w http.ResponseWriter, request *http.Request) {
    utils.RenderTemplate(w, request, "admin.html", nil)
}
