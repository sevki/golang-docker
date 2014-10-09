FROM golang

# Add source directory
ADD . /go/src/hello/

# Compile hello
RUN go install hello

# Fire up hello
ENTRYPOINT /go/bin/hello

# Expose the port
EXPOSE 8080
