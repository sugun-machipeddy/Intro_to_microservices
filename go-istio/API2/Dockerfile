FROM iron/go:dev
EXPOSE 8080
WORKDIR /app
ADD . /app
RUN cd /app; go build -o myapp2; 
ENTRYPOINT ["./myapp2"]

