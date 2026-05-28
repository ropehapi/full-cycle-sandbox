FROM golang:1.26.3

WORKDIR /app

COPY . .

RUN go build -o main

CMD ["./main"]