package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Gateway struct {
	name string
	port string
}

func getConfig() Gateway {
	config := Gateway {
		name: "Content Gateway",
		port: "4461",
	}
	return config
}

func handler(w http.ResponseWriter, r *http.Request) {
	content := "<div><h1>"+ getConfig().name +"</h1></div>"
	fmt.Fprint(w, content)
}

func main() {
	config := getConfig()
	port := ":" + config.port

	r := mux.NewRouter()
	r.HandleFunc("/", handler)

	fmt.Println(config.name + "Listening on port " + port)

	err := http.ListenAndServe(port, r)

	if err != nil {
		log.Fatal("Listen and serve: ", err)
		return
	}

}
