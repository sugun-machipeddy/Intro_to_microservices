FROM golang
EXPOSE 3000
WORKDIR /go/src/github.com/sugun/grpc-go/client
ADD . /go/src/github.com/sugun/grpc-go/client
RUN cd /go/src/github.com/sugun/grpc-go/client
RUN go get -v ./...
RUN go build -o myapp 
ENTRYPOINT ["./myapp"]
