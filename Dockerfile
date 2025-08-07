# ---------- STAGE 1: Build frontend (Vite) ----------
FROM node:20-alpine AS frontend

WORKDIR /frontend

COPY frontend/package.json frontend/package-lock.json* ./
RUN npm install

COPY frontend .

RUN npm run build


# ---------- STAGE 2: Build backend (Go) ----------
FROM golang:1.23-alpine AS backend

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

RUN apk add --no-cache git

COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend .
COPY backend/migrations ./migrations

RUN go build -o finance-api ./cmd/server


# ---------- STAGE 3: Final image ----------
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy Go binary and migrations
COPY --from=backend /app/finance-api .
COPY --from=backend /app/migrations ./migrations

# Copy frontend static files
COPY --from=frontend /frontend/dist ./frontend/dist

# Serve static files using Go
# (Assumes your Go server uses ./frontend/dist for static content)
EXPOSE 8080

CMD ["./finance-api"]
