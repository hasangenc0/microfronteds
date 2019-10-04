package types

type HTTPMethod string

type Gateway struct {
	Name string
	Host string
	Port string
	Method string
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

func (gateway *Gateway) GetHTTPMethod() string {
	switch gateway.Method{
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
		panic(gateway.Method + " is not a type of http method.")
	}
}

type Page struct {
	Content string
	Name string
}
