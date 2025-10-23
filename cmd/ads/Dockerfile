FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY cmd/ads/.env /app/.env


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /ads ./cmd/ads

FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=builder /ads .

EXPOSE 8081

ENTRYPOINT ["/app/ads"]
