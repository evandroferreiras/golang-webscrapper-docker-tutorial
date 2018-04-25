FROM golang:onbuild

RUN go get golang.org/x/tour/gotour
RUN go get github.com/go-redis/redis
RUN go get github.com/pilu/fresh

EXPOSE 8080