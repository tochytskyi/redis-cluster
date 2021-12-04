FROM golang:1.16-alpine

ENV HTTP_PORT=5000

WORKDIR /go/src/app

COPY ./ ./

RUN apk add --no-cache git

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o /app ./

EXPOSE $HTTP_PORT

ENTRYPOINT ["/app"]

