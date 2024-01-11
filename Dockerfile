FROM golang:1.19-alpine as builder

WORKDIR /app

COPY . .

RUN go mod init test && go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o main

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/main /main

EXPOSE 8080

CMD ["/main"]