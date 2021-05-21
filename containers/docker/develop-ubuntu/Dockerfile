FROM ubuntu:xenial

RUN \
  apt-get update && apt-get upgrade -q -y && \
  apt-get install -y --no-install-recommends golang git make gcc libc-dev ca-certificates && \
  git clone --depth 1 https://github.com/yunhailanuxgk/go-jinbao && \
  (cd JINBAO && make geth) && \
  cp JINBAO/build/bin/geth /geth && \
  apt-get remove -y golang git make gcc libc-dev && apt autoremove -y && apt-get clean && \
  rm -rf /JINBAO

EXPOSE 8545
EXPOSE 30303

ENTRYPOINT ["/geth"]
