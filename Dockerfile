FROM golang:1.20.2-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

EXPOSE ${HTTP_PORT}

RUN go build -o bin/app ./cmd/app/main.go

CMD ["bin/app", "-d"]