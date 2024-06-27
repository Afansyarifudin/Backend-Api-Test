# FROM golang:1.22 AS builder
# WORKDIR /app
# COPY . .
# RUN CGO_ENABLED=0 GOOS=linux go build .


# FROM alpine:latest
# WORKDIR /app
# COPY --from=builder /app/backend-api .
# CMD ["./backend-api"]

FROM golang:1.22-alpine

COPY . /app

WORKDIR /app

RUN go mod tidy

RUN go build -o backend-api .

EXPOSE 8080

CMD ["/app/backend-api"]