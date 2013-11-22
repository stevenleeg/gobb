package controllers

import (
    "net/http"
    "sirjtaa/utils"
)

func Admin(w http.ResponseWriter, request *http.Request) {
    utils.RenderTemplate(w, request, "admin.html", nil)
}
