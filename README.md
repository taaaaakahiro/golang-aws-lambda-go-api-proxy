# golang-lambda-apigateway

## build
```
$ GOOS=linux GOARCH=amd64 go build -o hello ./cmd/main.go && zip hello.zip hello
```