FROM golang:1.21
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o invoice-service main.go
CMD ["./invoice-service"]
