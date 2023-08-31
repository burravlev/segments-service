FROM golang:alpine

WORKDIR /server
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /main

CMD ["/main"]