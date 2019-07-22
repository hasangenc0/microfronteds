package collector

type HTTPMethod string

type Gateway struct {
	Name string
	Host string
	Port string
	Method string
}

type Page struct {
	Content string
	Name string
}

type App struct {
	Gateway []Gateway
	Page Page
}