FROM golang
EXPOSE 3000
WORKDIR /go/src/github.com/sugun/grpc-go/server
ADD . /go/src/github.com/sugun/grpc-go/server
RUN cd /go/src/github.com/sugun/grpc-go/server
RUN go get -v ./...
RUN go build -o myapp 
ENTRYPOINT ["./myapp"]
