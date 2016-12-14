FROM golang:latest

RUN go get -u github.com/markbates/buffalo/buffalo
RUN go install github.com/markbates/buffalo/buffalo
WORKDIR /go/src/github.com/gopheracademy/gcon

CMD buffalo dev