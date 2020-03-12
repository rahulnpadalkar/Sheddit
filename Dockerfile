FROM golang:1.13.4 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o sheddit .

FROM alpine:3.9
COPY .env .
COPY --from=builder /app/sheddit .
EXPOSE 7009
CMD ["./sheddit"]