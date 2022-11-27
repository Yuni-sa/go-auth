FROM golang:alpine AS builder

WORKDIR /build

COPY . .

RUN go build -o Auth

FROM alpine

WORKDIR /app

ENV PORT="4000"

COPY --from=builder /build/Auth /app/Auth

CMD ["./Auth"]