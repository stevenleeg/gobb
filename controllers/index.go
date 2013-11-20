package controllers

import (
    "net/http"
    "sirjtaa/utils"
)

func Index(w http.ResponseWriter, request *http.Request) {
    utils.RenderTemplate(w, request, "index.html", nil)
}
