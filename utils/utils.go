package utils

import (
    "html/template"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
)

var Store = sessions.NewCookieStore([]byte("83kjhsd98w3kjhwdfsdfw3"))

func RenderTemplate(
	out http.ResponseWriter,
	r *http.Request,
	tpl_file string,
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

    tpl, err := template.ParseFiles("templates/base.html", "templates/" + tpl_file)
    if err != nil {
        FatalError(err, "Template error")
    }
    tpl.Execute(out, send)
}

func FatalError(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
