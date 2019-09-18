package types

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
