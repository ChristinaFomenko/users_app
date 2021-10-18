FROM golang:1.16.5-buster

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o users_app ./cmd/users_app/main.go

CMD ["./users_app"]