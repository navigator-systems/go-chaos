FROM golang:1.26-alpine AS builder
WORKDIR /go-chaos
COPY . . 
RUN go build -ldflags="-s -w" -o go-chaos .
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go-chaos/go-chaos /bin/go-chaos
ENTRYPOINT [ "/bin/go-chaos" ]

