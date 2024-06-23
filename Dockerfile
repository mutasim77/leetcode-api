FROM golang:1.20-alpine AS builder

WORKDIR /app

# Add go.sum if needed
COPY go.mod  ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o leet-api ./cmd/server

FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/leet-api .

EXPOSE 8080

CMD ["./leet-api"]