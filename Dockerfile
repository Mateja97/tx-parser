FROM golang:1.23.4 AS build
WORKDIR /build

COPY go.mod go.sum ./
RUN go mod tidy
RUN go mod vendor

COPY . .

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o app


FROM golang:1.23.4

WORKDIR /app

COPY --from=build /build/app .

EXPOSE 8080
ENTRYPOINT ["/app/app"]