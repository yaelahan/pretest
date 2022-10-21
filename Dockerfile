FROM golang:1.19.2-alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o binary

FROM alpine
WORKDIR /app
COPY --from=builder /app/binary /app/

ENTRYPOINT ["/app/binary"]