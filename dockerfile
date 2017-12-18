FROM alpine:latest
MAINTAINER "github.com/lflxp"

ADD server /bin/server
RUN chmod +x /bin/server
EXPOSE 8972
ENTRYPOINT ["/bin/server"]
