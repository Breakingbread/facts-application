FROM golang:alpine3.18 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /facts-microservice .

FROM alpine:3.18 AS final

WORKDIR /app

RUN apk add --no-cache curl \
    && rm -rf /var/cache/apk/*

COPY --from=builder /facts-microservice .

EXPOSE 8080

ENTRYPOINT ["./facts-microservice"]


