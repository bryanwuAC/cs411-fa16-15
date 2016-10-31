package mvc

import (
	"html/template"
	"log"
	"net/http"
)

type loginController struct {
}

func (this *loginController) IndexAction(w http.ResponseWriter, r *http.Request) {
	log.Println("in the Login controller! function Executed for no reason")

	t, err := template.ParseFiles("web/admin.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)

}
