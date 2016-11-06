package main

import (
	"log"
	"net/http"
	"./mvc"
)

func main() {
	log.Println("main")
	http.Handle("/", http.FileServer(http.Dir("web")))
	http.HandleFunc("/ajax/", mvc.AjaxHandler)
	//http.ListenAndServe("fa16-cs411-15.cs.illinois.edu:8888", nil)
	http.ListenAndServe("localhost:8888", nil)

}
