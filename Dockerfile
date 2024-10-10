FROM golang:1.23-alpine

WORKDIR /app

# Copy files from the 'service' directory
COPY service/go.mod ./
COPY service/go.sum ./
RUN go mod download

COPY service/ . 

RUN go build -o main .

CMD ["./main"]