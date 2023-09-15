FROM golang:1.21.1-alpine AS build

WORKDIR /app

COPY go.mod go.sum .

RUN go mod download

COPY server.go .
COPY ldap/ ldap/
COPY config/ config/

RUN go build server.go

FROM alpine:3.18.3
LABEL org.opencontainers.image.title="Self Service Password"
LABEL org.opencontainers.image.description="Allows users to change there password on LDAP-Server"
LABEL org.opencontainers.image.authors="philskat"
LABEL org.opencontainers.image.source="https://github.com/philskat/self-service-password"
LABEL org.opencontainers.image.version="0.0.1"

WORKDIR /app

COPY --from=build /app/server .
COPY public/ public/

EXPOSE 8080

CMD ["./server"]
