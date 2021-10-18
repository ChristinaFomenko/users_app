FROM golang:1.16.5-buster

RUN go version
ENV GOPATH=/

COPY ./ ./

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# make wait_for-postgres.sh executable
#RUN chmod +x wait-for-postgres.sh

# build go app
RUN go mod download
RUN go build -o users_app ./cmd/users_app/main.go

CMD ["./users_app"]