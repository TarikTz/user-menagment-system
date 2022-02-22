FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 3010

CMD ["./wait-for.sh" , "mysql:3306" , "--timeout=300" , "--" , "./main"]