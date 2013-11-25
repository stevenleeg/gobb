package utils

import (
	"github.com/gorilla/sessions"
	"net/http"
	"github.com/stevenleeg/gobb/models"
	"github.com/stevenleeg/gobb/config"
)

var Store *sessions.CookieStore

func GetCookieStore(r *http.Request) *sessions.CookieStore {
    if Store == nil {
        cookie_key, _ := config.Config.GetString("gobb", "cookie_key")
        Store = sessions.NewCookieStore([]byte(cookie_key))
    }

    return Store
}

// TODO: Cache this!
func GetCurrentUser(r *http.Request) *models.User {
	session, _ := GetCookieStore(r).Get(r, "sirsid")

	if session.Values["username"] == nil || session.Values["password"] == nil {
		return nil
	}
	err, current_user := models.AuthenticateUser(session.Values["username"].(string), session.Values["password"].(string))

	if err != nil {
		return nil
	}
	return current_user
}
