FROM golang:alpine as builder
ENV BACKEND_HOST=localhost \
    BACKEND_PORT=8001

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o bin main.go

FROM scratch
WORKDIR /app
COPY --from=builder /app/bin .
EXPOSE 8000
CMD ["/app/bin"]