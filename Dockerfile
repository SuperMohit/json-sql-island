FROM golang:alpine AS builder
WORKDIR /go/src/app
COPY . .
RUN go build -o goserver .

FROM alpine
WORKDIR /app
COPY --from=builder /go/src/app/goserver /app/

CMD ["./goserver"]

