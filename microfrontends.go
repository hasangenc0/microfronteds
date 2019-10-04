package microfrontends

import (
	"fmt"
	"github.com/hasangenc0/microfrontends/pkg/client"
	"github.com/hasangenc0/microfrontends/pkg/types"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Gateway = types.Gateway
type Page = types.Page

type App struct {
	Gateway []Gateway
	Page Page
	Response http.ResponseWriter
}

func getUrl(host string, port string) string {
	return host + ":" + port
}

func (app *App) setHeaders() {
	app.Response.Header().Set("Transfer-Encoding", "chunked")
	//w.Header().Set("X-Content-Type-Options", "nosniff")
}

func (app *App) initialize() {
	flusher, ok := app.Response.(http.Flusher)

	if !ok {
		panic("expected http.ResponseWriter to be an http.Flusher")
	}

	tmpl, err := template.New(app.Page.Name).Parse(app.Page.Content)

	if err != nil {
		panic("An Error occured when parsing html")
		return
	}

	err = tmpl.Execute(app.Response, "")

	if err != nil {
		panic("Error in Template.Execute")
	}

	flusher.Flush()
}

func (app *App) sendChunk(gateway Gateway, ch chan http.Flusher) {
	var flusher, ok = app.Response.(http.Flusher)
	if !ok {
		panic("expected http.ResponseWriter to be an http.Flusher")
	}

	_client := &http.Client{}
	req, err := http.NewRequest(gateway.GetHTTPMethod(), getUrl(gateway.Host, gateway.Port), nil)
	if err != nil {
		panic(err)
	}
	resp, err := _client.Do(req)
	if err != nil {
		ch <- nil
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		bodyString := string(bodyBytes)

		chunk := client.GetView(gateway.Name, bodyString)

		fmt.Fprintf(app.Response, chunk)
	}

	ch <- flusher
}

func (app *App) finish() {
	flusher, ok := app.Response.(http.Flusher)

	if !ok {
		panic("expected http.ResponseWriter to be an http.Flusher")
	}

	_, err := app.Response.Write([]byte(""))

	if err != nil {
		panic("expected http.ResponseWriter to be an http.Flusher")

	}

	flusher.Flush()

}

func (app *App) Init() {
	app.setHeaders()

	app.initialize()

	var flusher = make(chan http.Flusher)

	for _, gateway := range app.Gateway {
		go app.sendChunk(gateway, flusher)
	}

	for range app.Gateway {
		flusher, ok := <-flusher
		if !ok {
			panic("expected http.ResponseWriter to be an http.Flusher")
		}
		if flusher != nil {
			flusher.Flush()
		}
	}

	app.finish()
}
