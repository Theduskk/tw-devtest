# syntax=docker/dockerfile:1

FROM golang:1.23.4

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./
COPY apis/ ./apis/
COPY structs/ ./structs/

RUN CGO_ENABLED=0 GOOS=linux go build -o /tw-devtask

EXPOSE 3000

CMD ["/tw-devtask"]

