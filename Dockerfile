FROM golang:1.22rc2-alpine3.19 as builder

WORKDIR /app

COPY . .

RUN go build main.go

# FROM alpine
FROM scratch

COPY --from=builder /app /app

CMD ["./app/main"]