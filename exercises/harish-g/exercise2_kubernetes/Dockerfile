FROM golang:1.13 AS goapp
WORKDIR /app
COPY hello.go .
RUN go build -o hello .
EXPOSE 8080
ENTRYPOINT ["./hello"]

