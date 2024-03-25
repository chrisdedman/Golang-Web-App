FROM debian:stretch-slim

WORKDIR /

RUN apt-get update && apt-get install -y ca-certificates

ADD bin /bin/

CMD ["/bin/sh"]
