# syntax=docker/dockerfile:1

FROM golang:1.21.0


WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

EXPOSE 8081
RUN go build -v -o /usr/local/bin/app ./...

CMD ["app"]