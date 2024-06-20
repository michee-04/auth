FROM golang:1.21.10

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download


COPY . .

RUN  CGO_ENABLED=0 GOOS=linux go build -o /authapp

EXPOSE 8080


CMD ["/authapp"]