#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /src/app
COPY . .
RUN go get -d -v ./...
RUN go build -o /bin/app -v ./...

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /bin/app /app
ENTRYPOINT /app
EXPOSE 8090
