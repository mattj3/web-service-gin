# syntax=docker/dockerfile:1

# specify the base image to  be used for the application, alpine or ubuntu
FROM golang:1.22-alpine

# create a working directory inside the image
WORKDIR /app

# copy Go modules and dependencies to image
COPY go.mod ./

RUN go mod tidy

# download Go modules and dependencies
RUN go mod download

# get github required dependencies (not sure why the above didn't work)
RUN go get github.com/gin-gonic/gin

# copy directory files i.e all files ending with .go
COPY *.go ./

# compile application
RUN go build -o /web-service-gin

# tells Docker that the container listens on specified network ports at runtime
EXPOSE 8080

# command to be used to execute when the image is used to start a container
CMD [ "/web-service-gin" ]