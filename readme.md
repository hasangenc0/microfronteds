# Microfrontends

This project is under development
```
go get github.com/hasangenc0/microfrontends
```

## Usage
Look at the [examples](./examples) folder

##### Running Examples
```
cd examples
sh run.sh
```

##### Define your microservice gateways
```go
gateways := []microfrontends.Gateway{
    {
        Name: "header",
        Host: "http://localhost",
        Port: "4462",
        Method: "GET",
    },
    {
        Name: "footer",
        Host: "http://localhost",
        Port: "4463",
        Method: "GET",
    },
    {
        Name: "content",
        Host: "http://localhost",
        Port: "4461",
        Method: "GET",
    },
}
```

##### Define your default page
```go
page := microfrontends.Page{
    Name: "App",
    Content: `
        <html>
        <body>
            <chunk name="header"></chunk>
            <chunk name="content"></chunk>
            <chunk name="footer"></chunk>
        </body>
        </html>
    `,
}
```

##### Create microfrontends app and run it
```go
app := microfrontends.App{
    gateways,
    page,
    w,
}

app.Init()
```

## Author
[Hasan Genc](https://www.linkedin.com/in/hasangenc0/) 
