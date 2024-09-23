FROM golang:1.22-alpine AS build

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o loan-management-system ./cmd/http

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=build /app/loan-management-system .

EXPOSE 9002

CMD ["./loan-management-system"]