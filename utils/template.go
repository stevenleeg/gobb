package utils

import (
	"github.com/stevenleeg/gobb/config"
	"github.com/stevenleeg/gobb/models"
    "github.com/russross/blackfriday"
	"html/template"
	"fmt"
    "time"
	"net/http"
)

func tplAdd(first, second int) int {
    return first + second
}

func tplParseMarkdown(input string) template.HTML {
    byte_slice := []byte(input)
    return template.HTML(string(blackfriday.MarkdownCommon(byte_slice)))
}

func tplGetCurrentUser(r *http.Request) func() *models.User {
    return func() *models.User {
        return GetCurrentUser(r)
    }
}

func tplIsValidTime(in time.Time) bool {
    return in.Year() > 1
}

var default_funcmap = template.FuncMap{
    "TimeRelativeToNow": TimeRelativeToNow,
    "Add": tplAdd,
    "ParseMarkdown": tplParseMarkdown,
    "IsValidTime": tplIsValidTime,
}


func RenderTemplate(
	out http.ResponseWriter,
	r *http.Request,
	tpl_file string,
	context map[string]interface{},
    funcs template.FuncMap) {

	current_user := GetCurrentUser(r)
	site_name, _ := config.Config.GetString("gobb", "site_name")
    ga_tracking_id, _ := config.Config.GetString("googleanalytics", "tracking_id")
    ga_account, _ := config.Config.GetString("googleanalytics", "account")

	send := map[string]interface{}{
		"current_user": current_user,
		"request": r,
		"site_name": site_name,
        "ga_tracking_id": ga_tracking_id,
        "ga_account": ga_account,  
	}

	// Merge the global template variables with the local context
	for key, val := range context {
		send[key] = val
	}

    // Same with the function map
    func_map := default_funcmap
    func_map["GetCurrentUser"] = tplGetCurrentUser(r)
    for key, val := range funcs {
        func_map[key] = val
    }

	tpl, err := template.New("tpl").Funcs(func_map).ParseFiles("templates/base.html", "templates/"+tpl_file)
	if err != nil {
        fmt.Printf("[error] Could not parse template (%s)\n", err.Error())
	}
	tpl.ExecuteTemplate(out, tpl_file, send)
	tpl.ExecuteTemplate(out, "base.html", send)
}

