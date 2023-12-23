FROM golang:1.21.3-alpine3.18

WORKDIR /app

COPY . ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /dating-app-api

EXPOSE 8080

CMD ["/dating-app-api"]