FROM golang:1.13-alpine
WORKDIR /app
COPY . .
RUN go build -o ./bin/server
EXPOSE 8080
ENTRYPOINT ["./bin/server"]
