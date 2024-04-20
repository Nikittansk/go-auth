FROM golang:1.22.2-alpine3.19

RUN go install github.com/githubnemo/CompileDaemon@latest

WORKDIR /app

COPY ./ ./

RUN go mod tidy

ENTRYPOINT CompileDaemon --build="go build -o build/go-auth" -command="./build/go-auth" -build-dir=/app
# FROM golang:1.22.2-alpine3.19

# WORKDIR /app

# COPY . .

# RUN go get -d -v .

# RUN go build -o go-auth .

# EXPOSE 8000

# CMD [ "./go-auth" ]