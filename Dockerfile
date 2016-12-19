FROM ubuntu
ADD . /gcon
WORKDIR /gcon
EXPOSE 3000
ENV GO_ENV production
CMD ["/gcon/gcon"]

