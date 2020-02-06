FROM golang:1.13 AS goapp
RUN mkdir /app
ADD main.go /app
WORKDIR /app
RUN go build -o main .

FROM ubi8:latest
WORKDIR /app
COPY --from=goapp /app/main .
ENTRYPOINT ["./main"]

