FROM golang:1.23.1

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o app .

EXPOSE 80
CMD ["./app"]
