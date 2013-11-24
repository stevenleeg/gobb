package utils

import (
	"github.com/gorilla/sessions"
	"github.com/stevenleeg/gobb/config"
    "github.com/russross/blackfriday"
	"html/template"
	"log"
	"net/http"
)

var Store = sessions.NewCookieStore([]byte("83kjhsd98w3kjhwdfsdfw3"))

func tpl_add(first, second int) int {
    return first + second
}

func tpl_markdown(input string) template.HTML {
    byte_slice := []byte(input)
    return template.HTML(string(blackfriday.MarkdownCommon(byte_slice)))
}

func RenderTemplate(
	out http.ResponseWriter,
	r *http.Request,
	tpl_file string,
	context map[string]interface{}) {

	func_map := template.FuncMap{
		"TimeRelativeToNow": TimeRelativeToNow,
        "add": tpl_add,
        "markdown": tpl_markdown,
	}

	current_user := GetCurrentUser(r)
	site_name, _ := config.Config.GetString("gobb", "sitename")

	send := map[string]interface{}{
		"current_user": current_user,
		"request":      r,
		"site_name":    site_name,
	}

	// Merge the global template variables with the local context
	for key, val := range context {
		send[key] = val
	}

	tpl, err := template.New("tpl").Funcs(func_map).ParseFiles("templates/base.html", "templates/"+tpl_file)
	if err != nil {
		FatalError(err, "Template error")
	}
	tpl.ExecuteTemplate(out, tpl_file, send)
	tpl.ExecuteTemplate(out, "base.html", send)
}

func FatalError(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
