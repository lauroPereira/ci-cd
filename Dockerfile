FROM golang:1.14

COPY . .

RUN go build main.go

EXPOSE 80
ENTRYPOINT ["./main"]