FROM golang:1.24.6-alpine AS builder                                        
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN go build -o /hello-svirin

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /hello-svirin .
EXPOSE 30757
CMD ["./hello-svirin"]