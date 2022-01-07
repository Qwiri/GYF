FROM golang:1.17 AS builder

LABEL maintainer="darmiel <hi@d2a.io>"
LABEL org.opencontainers.image.source = "https://github.com/qwiri/GYF"

WORKDIR /usr/src/app
SHELL ["/bin/bash", "-o", "pipefail", "-c"]

# Install dependencies
# Thanks to @montanaflynn
# https://github.com/montanaflynn/golang-docker-cache
COPY go.mod go.sum ./
RUN go mod graph | awk '{if ($1 !~ "@") print $2}' | xargs go get

# Copy remaining source
COPY . .

# Build from sources
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o backend ./cmd/backend

FROM alpine:3.15
COPY --from=builder /usr/src/app/backend .

EXPOSE 8080

ENTRYPOINT [ "/backend" ]