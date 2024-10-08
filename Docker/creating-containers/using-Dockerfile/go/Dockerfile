from golang:1.22

WORKDIR /app

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build main.go

EXPOSE 8080

CMD ["./main"]
