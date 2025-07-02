FROM golang:1.24

WORKDIR /worker

COPY go.mod ./
RUN go mod download

COPY . .
RUN go build -v -o ./cmd/worker ./...

CMD [ "./cmd/worker" ]