FROM golang:1.14 AS gobuilder
WORKDIR /app
COPY go.mod go.sum main.go /app/
RUN CGO_ENABLED=0 go build -mod=readonly -o app

FROM ubuntu:18.04
WORKDIR /app
COPY --from=gobuilder /app/app /app/app
CMD ["/app/app"]
