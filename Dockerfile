FROM golang:1.22.2-alpine3.19

RUN go install github.com/githubnemo/CompileDaemon@latest

WORKDIR /app

COPY ./ ./

RUN go mod tidy

ENTRYPOINT CompileDaemon --build="go build -o build/go-auth" -command="./build/go-auth" -build-dir=/app