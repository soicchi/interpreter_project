FROM golang:1.22rc1-bullseye
WORKDIR /go
COPY go.mod go.sum ./
RUN go mod download
COPY . .
