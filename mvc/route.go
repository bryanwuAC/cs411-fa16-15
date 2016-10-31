	package mvc

import (
	"html/template"
	"log"
	"net/http"
	"reflect"
	"strings"
	// "gopkg.in/gotsunami/coquelicot.v1"
)


func AjaxHandler(w http.ResponseWriter, r *http.Request) {
	// log.Println("ajaxHandler")
	pathInfo := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(pathInfo, "/")
	log.Println(parts)
	var action = ""
	if len(parts) > 1 {
		action = strings.Title(parts[1]) + "Action"
	}
	log.Println(action)

	ajax := &ajaxController{}	//same as a constructor
	controller := reflect.ValueOf(ajax)
	method := controller.MethodByName(action)

	if !method.IsValid() {
		method = controller.MethodByName(strings.Title("index") + "Action")
	}
	requestValue := reflect.ValueOf(r)
	responseValue := reflect.ValueOf(w)
	method.Call([]reflect.Value{responseValue, requestValue}) //called ajaxController here
}


func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("notfoundHandler")
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/login/index", http.StatusFound)
	}

	t, err := template.ParseFiles("template/html/404.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
}
