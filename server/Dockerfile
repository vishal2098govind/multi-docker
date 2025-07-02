FROM golang:1.24

WORKDIR /server

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -v -o server ./...

CMD ["./server"]