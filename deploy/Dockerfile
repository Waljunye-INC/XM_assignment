FROM golang:1.24-alpine AS builder
LABEL authors="matvei_volkov"
RUN apk add --no-cache git ca-certificates

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/xm_assignment ./cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /bin/xm_assignment /app/xm_assignment

RUN chmod +x /app/xm_assignment

EXPOSE 8012

ENV DB_HOST=${DB_HOST}
ENV DB_PORT=${DB_PORT}
ENV DB_NAME=${DB_NAME}
ENV DB_USER=${DB_USER}
ENV DB_PASSWORD=${DB_PASSWORD}
ENV PUBLIC_API_PORT=${PUBLIC_API_PORT}
ENV JWT_KEY=${JWT_KEY}
ENV KAFKA_TOPIC=${KAFKA_TOPIC}
ENV BROKERS=${BROKERS}

CMD ["./xm_assignment", "-cfg-file=/app/.env"]
