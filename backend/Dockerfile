FROM golang:latest AS builder

WORKDIR /app
COPY ../backend /app/backend

WORKDIR /app/backend
RUN go mod download

WORKDIR /app/backend/cmd
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./sade-backend

FROM alpine:latest

RUN apk add --no-cache ffmpeg

WORKDIR /app
COPY --from=builder /app/backend /app/backend

WORKDIR /app/backend/cmd
RUN chmod +x ./sade-backend

EXPOSE 8080
CMD ["/app/backend/cmd/sade-backend"]