FROM golang:latest

WORKDIR /usr/src/app

# cache dependencies
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app .
EXPOSE 8080

CMD ["/usr/local/bin/app"]