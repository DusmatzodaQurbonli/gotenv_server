# --- STAGE 1: Builder ---
FROM golang:1.23.6 AS builder

WORKDIR /app

ARG VERSION=dev
ARG COMMIT=none

COPY go.mod go.sum ./
RUN go mod download

COPY . .

COPY configs/configs.json ./configs

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init

ENV CGO_ENABLED=0 GOOS=linux
RUN go build -trimpath \
  -ldflags="-s -w -X 'main.version=${VERSION}' -X 'main.commit=${COMMIT}'" \
  -o gotenv main.go

# --- STAGE 2: Minimal runtime ---
FROM ubuntu:22.04

ENV DEBIAN_FRONTEND=noninteractive

WORKDIR /app

COPY --from=builder /app/gotenv .
COPY --from=builder /app/docs ./docs
COPY --from=builder /app/configs ./configs

RUN apt-get update && \
    apt-get install -y ca-certificates curl tzdata && \
    rm -rf /var/lib/apt/lists/*

ENV TZ=Asia/Dushanbe

EXPOSE 4545

HEALTHCHECK --interval=30s --timeout=5s --start-period=10s --retries=3 \
  CMD curl --fail http://localhost:4545/ping || exit 1

ENTRYPOINT ["./gotenv"]
