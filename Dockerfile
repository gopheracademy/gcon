FROM alpine
ADD gcon /gcon/gcon
ADD templates/ /gcon/
ADD assets/ /gcon/
WORKDIR /gcon
ENTRYPOINT ["/gcon/gcon"]
