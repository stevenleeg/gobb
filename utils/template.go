package utils

import (
	"github.com/stevenleeg/gobb/config"
	"github.com/stevenleeg/gobb/models"
    "github.com/russross/blackfriday"
	"html/template"
	"fmt"
    "time"
	"net/http"
    "go/build"
    "path/filepath"
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

    stylesheet := "/static/style.css"
    if (current_user != nil) && current_user.StylesheetUrl.Valid && current_user.StylesheetUrl.String != "" {
    	stylesheet = current_user.StylesheetUrl.String
    }

	send := map[string]interface{}{
		"current_user": current_user,
		"request": r,
		"site_name": site_name,
        "ga_tracking_id": ga_tracking_id,
        "ga_account": ga_account,  
        "stylesheet": stylesheet,  
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

    // Get the base template path
    pkg, _ := build.Import("github.com/stevenleeg/gobb/gobb", ".", build.FindOnly)
    base_path := filepath.Join(pkg.SrcRoot, pkg.ImportPath, "templates/base.html")
    tpl_path := filepath.Join(pkg.SrcRoot, pkg.ImportPath, "templates/" + tpl_file)

	tpl, err := template.New("tpl").Funcs(func_map).ParseFiles(base_path, tpl_path)
	if err != nil {
        fmt.Printf("[error] Could not parse template (%s)\n", err.Error())
	}

    // Attempt to execute the template we're on
	err = tpl.ExecuteTemplate(out, tpl_file, send)
	if err != nil {
        fmt.Printf("[error] Could not parse template (%s)\n", err.Error())
	}

    // And now the base template
	err = tpl.ExecuteTemplate(out, "base.html", send)
	if err != nil {
        fmt.Printf("[error] Could not parse template (%s)\n", err.Error())
	}
}

