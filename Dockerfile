FROM golang:alpine AS builder
WORKDIR /app
COPY ./stakeholders-service .
EXPOSE 8082
ENTRYPOINT ["go", "run", "main.go"]