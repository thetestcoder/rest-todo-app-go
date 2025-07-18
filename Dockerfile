FROM golang:1.24.4-alpine AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o my-app ./cmd/rest

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/my-app ./
COPY --from=builder /app/configs ./configs

EXPOSE 8889

CMD ["./my-app"]