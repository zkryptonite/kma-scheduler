FROM golang:alpine3.14

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN go build

EXPOSE 5000

CMD ["./backend"]