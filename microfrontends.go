package microfrontends

import (
	"github.com/hasangenc0/microfrontends/pkg/client"
	"github.com/hasangenc0/microfrontends/pkg/collector"
	"html/template"
	"net/http"
	"runtime"
	"sync"
)

type Gateway = collector.Gateway
type Page = collector.Page
type App = collector.App

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Transfer-Encoding", "chunked")
	//w.Header().Set("X-Content-Type-Options", "nosniff")
}

func initialize(w http.ResponseWriter, page Page) {
	flusher, ok := w.(http.Flusher)

	if !ok {
		panic("expected http.ResponseWriter to be an http.Flusher")
	}

	tmpl, err := template.New(page.Name).Parse(page.Content)

	if err != nil {
		panic("An Error occured when parsing html")
		return
	}

	err = tmpl.Execute(w, "")

	if err != nil {
		panic("Error in Template.Execute")
	}

	flusher.Flush()
}

func sendChunk(w http.ResponseWriter, gateway Gateway, wg *sync.WaitGroup, ch chan http.Flusher) {
	var flusher, ok = w.(http.Flusher)
	if !ok {
		panic("expected http.ResponseWriter to be an http.Flusher")
	}

	chunk := client.GetView(gateway.Name, gateway.Content)

	tmpl, err := template.New(gateway.Name).Parse(chunk)

	if err != nil {
		panic("An Error occured when parsing html")
	}

	err = tmpl.Execute(w, "")

	if err != nil {
		panic(err)
	}

	//flusher.Flush()
	ch <- flusher
	wg.Done()
}

func finish(w http.ResponseWriter) {
	flusher, ok := w.(http.Flusher)

	if !ok {
		panic("expected http.ResponseWriter to be an http.Flusher")
	}

	_, err := w.Write([]byte(""))

	if err != nil {
		panic("expected http.ResponseWriter to be an http.Flusher")

	}

	flusher.Flush()

}

func Make(w http.ResponseWriter, app App) {
	setHeaders(w)

	var wg sync.WaitGroup

	initialize(w, app.Page)

	runtime.GOMAXPROCS(4)

	var flusher = make(chan http.Flusher)

	for _, gateway := range app.Gateway {
		wg.Add(1)
		go sendChunk(w, gateway, &wg, flusher)
	}

	for range app.Gateway {
		flusher, ok := <-flusher
		if !ok {
			panic("expected http.ResponseWriter to be an http.Flusher")
		}
		flusher.Flush()
	}

	wg.Wait()

	finish(w)
}
