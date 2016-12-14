FROM golang:1.7
ADD . /go/src/github.com/gopheracademy/gcon
WORKDIR /go/src/github.com/gopheracademy/qcon
ENTRYPOINT ["/go/src/github.com/gopheracademy/gcon/gcon"]
