FROM golang:1.24

WORKDIR "/server"

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -v -o ./cmd/server ./...

CMD ["./cmd/server"]