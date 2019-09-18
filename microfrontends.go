package microfrontends

import (
	"fmt"
	"github.com/hasangenc0/microfrontends/pkg/client"
	"github.com/hasangenc0/microfrontends/pkg/types"
	"html/template"
	"io/ioutil"
	"net/http"
	"runtime"
	"sync"
)

type Gateway = types.Gateway
type Page = types.Page

type App struct {
	Gateway []Gateway
	Page Page
	Response http.ResponseWriter
}

const (
	MethodGet     = "GET"
	MethodHead    = "HEAD"
	MethodPost    = "POST"
	MethodPut     = "PUT"
	MethodPatch   = "PATCH"
	MethodDelete  = "DELETE"
	MethodConnect = "CONNECT"
	MethodOptions = "OPTIONS"
	MethodTrace   = "TRACE"
)

func getMethod(method string) string {
	switch method {
	case MethodGet: return MethodGet
	case MethodHead: return MethodHead
	case MethodPost: return MethodPost
	case MethodPut: return MethodPut
	case MethodPatch: return MethodPatch
	case MethodDelete: return MethodDelete
	case MethodConnect: return MethodConnect
	case MethodOptions: return MethodOptions
	case MethodTrace: return MethodTrace
	default:
		panic(method + " is not a type of http method.")
	}
}

func getUrl(host string, port string) string {
	return host + ":" + port
}

func (app App) setHeaders() {
	app.Response.Header().Set("Transfer-Encoding", "chunked")
	//w.Header().Set("X-Content-Type-Options", "nosniff")
}

func (app App) initialize() {
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

func (app App) sendChunk(gateway Gateway, wg *sync.WaitGroup, ch chan http.Flusher) {
	var flusher, ok = app.Response.(http.Flusher)
	if !ok {
		panic("expected http.ResponseWriter to be an http.Flusher")
	}

	_client := &http.Client{}
	req, err := http.NewRequest(getMethod(gateway.Method), getUrl(gateway.Host, gateway.Port), nil)
	if err != nil {
		panic(err)
	}
	resp, err := _client.Do(req)
	if err != nil {
		ch <- nil
		wg.Done()
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
	wg.Done()
}

func (app App) finish() {
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

func (app App) Init() {
	app.setHeaders()

	var wg sync.WaitGroup

	app.initialize()

	runtime.GOMAXPROCS(4)

	var flusher = make(chan http.Flusher)

	for _, gateway := range app.Gateway {
		wg.Add(1)
		go app.sendChunk(gateway, &wg, flusher)
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

	wg.Wait()

	app.finish()
}
