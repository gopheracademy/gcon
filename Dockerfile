FROM ubuntu
ADD ./assets /gcon/assets
ADD ./grifts /gcon/grifts
ADD ./migrations /gcon/migrations
ADD ./templates /gcon/templates
ADD database.yml /gcon/database.yml
ADD gcon /gcon/gcon
RUN chmod +x /gcon/gcon
WORKDIR /gcon
EXPOSE 3000
ENV GO_ENV production
CMD ["/gcon/gcon"]
