
FROM --platform=linux/arm64 golang:1.24-alpine
WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download

COPY . .

COPY data data

RUN go build -o main .

EXPOSE 8000

CMD ["./main"]