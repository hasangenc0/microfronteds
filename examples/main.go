package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hasangenc0/microfrontends"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	gateways := []microfrontends.Gateway{
		{
			Name: "header",
			Content: "<div><h1>Header</h1></div>",
		},
		{
			Name: "content",
			Content: "<div><h1>Content</h1></div>",
		},
		{
			Name: "footer",
			Content: "<div><h1>Footer</h1></div>",
		},
	}

	page := microfrontends.Page{
		Name: "App",
		Content: `
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<title>Microfrontends Example</title>
			</head>
			<body>
				<chunk name="header"></chunk>
				<chunk name="content"></chunk>
				<chunk name="footer"></chunk>
			</body>
			</html>
		`,
	}

	app := microfrontends.App{
		gateways,
		page,
	}

	microfrontends.Make(w, app);
}

func main() {
	port := ":4446"

	r := mux.NewRouter()
	r.HandleFunc("/", handler)

	fmt.Println("Listening...")

	err := http.ListenAndServe(port, r)

	if err != nil {
		log.Fatal("Listen and serve: ", err)
		return
	}

}
