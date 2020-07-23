# Specify parent image to build from
FROM golang:latest
# Set working directory to run following commands from
WORKDIR /go/src/hello
COPY . .
# N.B. triple dots are recursive wildcards in Go
RUN go get -d -v ./...
RUN go install -v ./...
# Do the thing!
CMD ["hello"]

