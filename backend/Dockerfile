FROM golang:alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o bin main.go

FROM scratch
WORKDIR /app
COPY --from=builder /app/bin .
EXPOSE 8001
CMD ["/app/bin"]