FROM golang:latest AS builder

WORKDIR /app
COPY ../backend /app/backend

WORKDIR /app/backend
RUN go mod download

WORKDIR /app/backend/cmd
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./sade-backend

FROM alpine:latest

RUN apk add --no-cache ffmpeg bash

COPY ../wait-for-it.sh /usr/local/bin/wait-for-it.sh
RUN chmod +x /usr/local/bin/wait-for-it.sh

WORKDIR /app
COPY --from=builder /app/backend /app/backend

WORKDIR /app/backend/cmd
RUN chmod +x ./sade-backend

EXPOSE 8080
ENTRYPOINT ["wait-for-it.sh", "db:5433", "--"]
CMD ["/app/backend/cmd/sade-backend"]
