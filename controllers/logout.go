package controllers

import (
	"github.com/stevenleeg/gobb/utils"
	"net/http"
    "fmt"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := utils.Store.Get(r, "sirsid")
	session.Values["username"] = ""
	session.Values["password"] = ""

	err := session.Save(r, w)
	if err != nil {
        fmt.Printf("[error] Could not save session (%s)\n", err.Error)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
