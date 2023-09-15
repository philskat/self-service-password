FROM golang:1.21.1-alpine AS build

WORKDIR /app

COPY go.mod go.sum .

RUN go mod download

COPY server.go .
COPY ldap/ ldap/
COPY config/ config/

RUN go build server.go

FROM alpine:3.18.3

WORKDIR /app

COPY --from=build /app/server .
COPY public/ public/

EXPOSE 8080

CMD ["./server"]
