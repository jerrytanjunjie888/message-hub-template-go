# --- Build stage ---
FROM golang:1.13 AS builder
WORKDIR /app
ENV GOPROXY=https://goproxy.io,direct
# Copy the dependency definition
COPY go.mod .
COPY go.sum .
# Download dependencies
RUN go mod download
# Copy the source code
COPY . .
# Build for release
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

# --- Final stage ---
# FROM gcr.io/distroless/static:latest
FROM registry-cn.subsidia.org/distroless/static:latest
COPY --from=builder /app/message-hub-template-go /
# VOLUME [ "/tmp/logs/" ]
ENTRYPOINT ["/message-hub-template-go"]
EXPOSE 3006
LABEL Name=message-hub-template-go \
	Version=0.1