
FROM golang:1.22 as builder


WORKDIR /app


COPY . .


RUN go mod download


RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o torontotime .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/torontotime .


EXPOSE 8014


CMD ["./torontotime"]
