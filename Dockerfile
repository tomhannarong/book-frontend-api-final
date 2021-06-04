FROM golang:alpine

RUN apk add --no-cache git

WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

RUN go get -u github.com/joho/godotenv

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build server.go" --command=./server

EXPOSE 8080