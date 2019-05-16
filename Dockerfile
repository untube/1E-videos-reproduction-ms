FROM golang:latest

WORKDIR $GOPATH/src/VideoPlayer-ms
COPY . .    

RUN apt-get update
RUN apt-get install vim -y

RUN go get -d -v ./...
RUN go build

CMD ["./VideoPlayer-ms"] 
