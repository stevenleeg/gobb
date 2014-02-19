package controllers

import (
	"fmt"
	"github.com/stevenleeg/gobb/utils"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := utils.GetCookieStore(r).Get(r, "sirsid")
	delete(session.Values, "username")
	delete(session.Values, "password")

	err := session.Save(r, w)
	if err != nil {
		fmt.Printf("[error] Could not save session (%s)\n", err.Error)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
