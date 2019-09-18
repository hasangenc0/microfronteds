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
			Host: "http://localhost",
			Port: "4462",
			Method: "GET",
		},
		{
			Name: "footer",
			Host: "http://localhost",
			Port: "4463",
			Method: "GET",
		},
		{
			Name: "content",
			Host: "http://localhost",
			Port: "4461",
			Method: "GET",
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
		w,
	}

	app.Init()
}

func main() {
	port := ":4460"

	r := mux.NewRouter()
	r.HandleFunc("/", handler)

	fmt.Println("Listening...")

	err := http.ListenAndServe(port, r)

	if err != nil {
		log.Fatal("Listen and serve: ", err)
		return
	}

}
