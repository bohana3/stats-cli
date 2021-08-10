FROM golang:1.16-alpine

COPY . /app
WORKDIR /app

# Download all the dependencies
RUN go get -d -v ./...

RUN apk update && apk add bash


# Install the package
RUN go install -v ./...

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
ENTRYPOINT ["/go/bin/stats-cli"]
