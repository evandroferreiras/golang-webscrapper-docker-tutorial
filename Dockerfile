FROM golang:onbuild

# ADD . /go/src/golang-webscrapper-docker-tutorial


# ENTRYPOINT /go/bin/golang-webscrapper-docker-tutorial

# RUN go install golang-webscrapper-docker-tutorial

RUN go get golang.org/x/tour/gotour

EXPOSE 8080