package controllers

import (
    "net/http"
    "sirjtaa/utils"
)

func Logout(w http.ResponseWriter, r *http.Request) {
    session, _ := utils.Store.Get(r, "sirsid")
    session.Values["username"] = ""
    session.Values["password"] = ""

    err := session.Save(r, w)
    if err != nil {
        utils.FatalError(err, "Could not save session!")
    }

    http.Redirect(w, r, "/", http.StatusFound)
}
