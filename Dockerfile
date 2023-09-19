# Build stage
FROM golang:1.21.1-alpine3.18 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .

RUN go build -o main main.go

# Run stage
FROM alpine:3.18.3
WORKDIR /app

COPY --from=builder /app/main .
COPY app.env .

EXPOSE 3000
CMD [ "/app/main" ]