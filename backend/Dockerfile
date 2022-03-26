FROM golang:1.17 AS builder

WORKDIR .
COPY . .
RUN mkdir /core
ENV GOPATH=/core
WORKDIR /app


# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . ./
RUN go build -o ./main/ ./...

CMD ["/app/main/virtualwhiteboard"]