package controllers

import (
    "net/http"
    "fmt"
    "github.com/stevenleeg/gobb/utils"
    "github.com/stevenleeg/gobb/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
    if utils.GetCurrentUser(r) != nil {
        http.Redirect(w, r, "/", http.StatusFound)
        return
    }

    if r.Method == "POST" {
        username := r.FormValue("username")
        password := r.FormValue("password")
        confirm  := r.FormValue("password2")

        var error string
        if(password != confirm) {
            error = "Passwords don't match"
        }

        if error != "" {
            utils.RenderTemplate(w, r, "register.html", map[string]interface{} {
                "error": error,
            })
            return
        }

        // We're good, let's make it
        db_map := models.GetDbSession()
        user := models.NewUser(username, password)
        err := db_map.Insert(user)
        
        if err != nil {
            fmt.Printf("[error] Could not insert user (%s)\n", err.Error())
        }
    }

    utils.RenderTemplate(w, r, "register.html", nil)
}
