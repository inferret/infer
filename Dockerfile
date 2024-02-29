FROM golang:latest
WORKDIR /app
COPY ./ /app
RUN go build ./cmd/infer/infer.go
CMD ["/app/infer"]