go run ./content_gateway/content_gateway.go &
go run ./header_gateway/header_gateway.go &
go run ./footer_gateway/footer_gateway.go &
go run main.go

wait