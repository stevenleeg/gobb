package controllers

import (
    "net/http"
    "fmt"
    "sirjtaa/utils"
    "sirjtaa/models"
)

func Index(w http.ResponseWriter, request *http.Request) {
    db := models.GetDbSession()
    var boards []models.Board
    _, err := db.Select(&boards, "SELECT * FROM boards")
    if err != nil {
        fmt.Printf("[error] Could not get boards (%s)\n", err.Error())
    }

    utils.RenderTemplate(w, request, "index.html", map[string]interface{} {
        "boards": boards,
    })
}
