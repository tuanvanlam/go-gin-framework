FROM golang:alpine

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 5000

RUN GOOS=linux go build -o server server.go

RUN find . -name "*.go" -type f -delete

EXPOSE $PORT

CMD ["./server"]

