FROM golang:latest

WORKDIR $GOPATH/src/video-reproduction-ms
COPY . .

RUN go get -d -v ./...
RUN go build

CMD ["./video-reproduction-ms"] 

EXPOSE 3002
