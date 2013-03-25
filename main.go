package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"web"
)

func HelloWorld(wr http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("./in_3.html")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = t.Execute(wr, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	web.Get("/(.*)", HelloWorld)
	web.Run("0.0.0.0:9999")
}
