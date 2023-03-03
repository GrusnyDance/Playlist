FROM golang:1.18 AS builder
WORKDIR /build

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/main.go

FROM golang AS server

WORKDIR /hflabs-doc

COPY --from=builder /build/app .
COPY --from=builder /build/.env .

CMD ["/hflabs-doc/./app"]