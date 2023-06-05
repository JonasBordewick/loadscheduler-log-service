# syntax=docker/dockerfile:1
FROM golang:alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /docker-loadscheduler-log-service
EXPOSE 49497
WORKDIR /config
CMD ["/docker-loadscheduler-log-service"]