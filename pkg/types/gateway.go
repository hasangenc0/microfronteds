package types

type Gateway struct {
	Name string
	Host string
	Port string
	Method string
}

func (gateway Gateway) GetHTTPMethod() string{
	switch gateway.Method {
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

func (gateway Gateway) GetUrl() string {
	return gateway.Host + ":" + gateway.Port
}
