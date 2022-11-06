FROM golang:1.19-alpine as builder
LABEL maintainer = "Rshezarr's Little Api"
WORKDIR /app
COPY . .
RUN go install github.com/pressly/goose/v3/cmd/goose@latest && go build -o api_main ./main.go

FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app .
RUN apk add bash
EXPOSE 8080
CMD ["/app/api_main"]