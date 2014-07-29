package utils

import (
	"fmt"
	"github.com/russross/blackfriday"
	"github.com/stevenleeg/gobb/config"
	"github.com/stevenleeg/gobb/models"
	"go/build"
	"html/template"
	"io/ioutil"
	"net/http"
	"path"
	"path/filepath"
	"strings"
	"time"
)

// Returns a list of all available themes
func ListTemplates() []string {
	names := []string{"default"}

	static_path, _ := config.Config.GetString("gobb", "base_path")
	files, _ := ioutil.ReadDir(path.Join(static_path, "templates"))

	for _, f := range files {
		if !f.IsDir() {
			continue
		}
		names = append(names, f.Name())
	}

	return names
}

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

func tplGetStringSetting(key string) string {
	val, _ := models.GetStringSetting(key)
	return val
}

func tplIsValidTime(in time.Time) bool {
	return in.Year() > 1
}

func tplParseFaviconType(url string) string {
	split := strings.Split(url, ".")
	if len(split) == 0 {
		return ""
	}
	return split[len(split)-1]
}

var default_funcmap = template.FuncMap{
	"TimeRelativeToNow": TimeRelativeToNow,
	"Add":               tplAdd,
	"ParseMarkdown":     tplParseMarkdown,
	"IsValidTime":       tplIsValidTime,
	"GetStringSetting":  tplGetStringSetting,
	"ParseFaviconType":  tplParseFaviconType,
}

func RenderTemplate(
	out http.ResponseWriter,
	r *http.Request,
	tpl_file string,
	context map[string]interface{},
	funcs template.FuncMap) {

	current_user := GetCurrentUser(r)
	site_name, _ := config.Config.GetString("gobb", "site_name")
	base_url, _ := config.Config.GetString("gobb", "base_url")
	ga_tracking_id, _ := config.Config.GetString("googleanalytics", "tracking_id")
	ga_account, _ := config.Config.GetString("googleanalytics", "account")

	stylesheet := ""
	if (current_user != nil) && current_user.StylesheetUrl.Valid && current_user.StylesheetUrl.String != "" {
		stylesheet = current_user.StylesheetUrl.String
	} else if current_user == nil || !current_user.StylesheetUrl.Valid || current_user.StylesheetUrl.String == "" {
		global_theme, _ := models.GetStringSetting("theme_stylesheet")
		if global_theme != "" {
			stylesheet = global_theme
		}
	}

	favicon_url, _ := models.GetStringSetting("favicon_url")

	send := map[string]interface{}{
		"current_user":   current_user,
		"request":        r,
		"site_name":      site_name,
		"ga_tracking_id": ga_tracking_id,
		"ga_account":     ga_account,
		"stylesheet":     stylesheet,
		"favicon_url":    favicon_url,
		"base_url":		  base_url,
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
	selected_template, _ := models.GetStringSetting("template")
	var base_path string
	if selected_template == "default" {
		pkg, _ := build.Import("github.com/stevenleeg/gobb/gobb", ".", build.FindOnly)
		base_path = filepath.Join(pkg.SrcRoot, pkg.ImportPath, "../templates/")
	} else {
		base_path, _ = config.Config.GetString("gobb", "base_path")
		base_path = filepath.Join(base_path, "templates", selected_template)
	}

	base_tpl := filepath.Join(base_path, "base.html")
	rend_tpl := filepath.Join(base_path, tpl_file)

	tpl, err := template.New("tpl").Funcs(func_map).ParseFiles(base_tpl, rend_tpl)
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
