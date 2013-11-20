package utils

import (
    "net/http"
    "sirjtaa/models"
)

// TODO: Cache this!
func GetCurrentUser(r *http.Request) *models.User {
	session, _ := Store.Get(r, "sirsid")

    err, current_user := models.AuthenticateUser(session.Values["username"].(string), session.Values["password"].(string))

	if err != nil {
		return nil
	}
	return current_user
}
