ARG BUILD_STAGE_BASE_IMAGE=golang:1.23.4
ARG DISTRIBUTION_STAGE_BASE_IMAGE=debian:bookworm-20241223-slim

# Build stage
FROM ${BUILD_STAGE_BASE_IMAGE} AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-redis-sample ./cmd/main.go


# Distribution stage
FROM ${DISTRIBUTION_STAGE_BASE_IMAGE}

WORKDIR /root/

RUN apt-get update && apt-get install -y --no-install-recommends \
  ca-certificates \
  && rm -rf /var/lib/apt/lists/*

USER nobody

COPY --from=builder /go-redis-sample .

CMD ["./go-redis-sample"] 