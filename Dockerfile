FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /fileshare


FROM alpine:latest

WORKDIR /
COPY --from=builder /fileshare /fileshare

EXPOSE 8080

ENTRYPOINT ["/fileshare"]