FROM golang:alpine AS builder
RUN mkdir /build
COPY . /build
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o axon cmd/axon/main.go

# copy only the binary into the container
FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN mkdir /app
WORKDIR /app
COPY --from=builder /build/axon .
CMD ["/app/axon"]


