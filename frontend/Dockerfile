FROM node:latest AS builder

WORKDIR /app
COPY ../frontend /app/frontend

WORKDIR /app/frontend
RUN npm install && npm run build

FROM node:latest
WORKDIR /app
COPY --from=builder /app/frontend /app/frontend

WORKDIR /app/frontend
RUN npm install --production

EXPOSE 3000
CMD ["node", "build"]
