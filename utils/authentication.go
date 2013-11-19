package utils

import (
    "net/http"
    "sirjtaa/models"
)

func GetCurrentUser(r *http.Request) *models.User {
	session, _ := Store.Get(r, "sirsid")

    err, current_user := models.AuthenticateUser(session.Values["username"].(string), session.Values["password"].(string))
	FatalError(err, "could not select from db")

	if current_user.Id == 0 {
		return nil
	}
	return current_user
}
