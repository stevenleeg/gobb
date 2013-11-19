package utils

import (
	"fmt"
	"github.com/hoisie/mustache"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
)

var Store = sessions.NewCookieStore([]byte("83kjhsd98w3kjhwdfsdfw3"))

func RenderTemplate(
	out http.ResponseWriter,
	r *http.Request,
	template string,
	context map[string]interface{}) {

    current_user := GetCurrentUser(r)
	send := map[string]interface{}{
		"current_user": current_user,
        "request": r,
	}

	// Merge the global template variables with the local context
	for key, val := range context {
		send[key] = val
	}

	rendered := mustache.RenderFileInLayout("templates/"+template, "templates/base.mustache", send)

	fmt.Fprint(out, rendered)
}

func FatalError(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
