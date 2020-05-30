FROM golang:1.14.2-alpine as builder

WORKDIR /go/src/app
ADD . /go/src/app

RUN go get -d -v ./...

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/app

FROM gcr.io/distroless/base
COPY --from=builder /go/bin/app /
CMD ["/app"]