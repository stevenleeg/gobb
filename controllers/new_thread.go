package controllers

import (
    "net/http"
    "sirjtaa/utils"
    //"sirjtaa/models"
)

func NewThread(w http.ResponseWriter, request *http.Request) {
    utils.RenderTemplate(w, request, "new_thread.html", nil)
}
