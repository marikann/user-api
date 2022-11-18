# stage-1: binary builder
FROM golang:1.17.1-alpine AS builder

## cache deps
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

## build binary
WORKDIR /app
COPY . .
RUN swag init -g main.go
RUN go build -o user-api .

# stage-2: image builder
FROM alpine
WORKDIR /build
COPY --from=builder /app/user-api .

EXPOSE 8080

# run
ENTRYPOINT [ "/build/user-api" ]