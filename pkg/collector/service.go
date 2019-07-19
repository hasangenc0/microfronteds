package collector

// Service is the model for microservice endpoints
type Service struct {
	host string
	port string
}

type Gateway struct {
	Name string
	Content string
}

type Page struct {
	Content string
	Name string
}

type App struct {
	Gateway []Gateway
	Page Page
}