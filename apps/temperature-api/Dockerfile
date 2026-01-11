FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod init temperature-api || true
RUN go mod tidy
RUN go build -o temperature-api main.go

FROM alpine:3.18

WORKDIR /app
COPY --from=builder /app/temperature-api .

EXPOSE 8081
CMD ["./temperature-api"]