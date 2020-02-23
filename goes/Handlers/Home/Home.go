package Home

import (
	"fmt"
	Handlers "github.com/zendrulat1/goes/Handlers/Headers"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template
var err error

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

//get handler
func Home(w http.ResponseWriter, r *http.Request) {
	Handlers.Header(w, r)
	switch r.Method {
	case "GET":
		err := tpl.ExecuteTemplate(w, "home.html", nil)
		if err != nil {
			log.Fatalln("template didn't execute: ", err)
		}

	case "POST":
		fmt.Println(r.Header.Get("Origin"))
		err := tpl.ExecuteTemplate(w, "home.html", nil)
		if err != nil {
			log.Fatalln("template didn't execute: ", err)
		}

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}

}