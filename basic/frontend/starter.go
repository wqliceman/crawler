package main

import (
	"github.com/wqliceman/crawler/basic/frontend/controller"
	"net/http"
)


const basePath = "/Users/iceman/Develops/dev/go/src/github.com/wqliceman/crawler/basic/frontend/"
func main() {
	http.Handle("/", http.FileServer(
		http.Dir(basePath + "view")))
	http.Handle("/search",
		controller.CreateSearchResultHandler(
			basePath + "view/template.html"))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
