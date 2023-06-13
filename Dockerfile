FROM ubuntu:focal as build

ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update && apt-get install -y curl make

RUN curl -L https://go.dev/dl/go1.18.10.linux-amd64.tar.gz |tar -xz -C /usr/local
ENV PATH=$PATH:/usr/local/go/bin

COPY . /g
WORKDIR /g
RUN make clean && make

FROM ubuntu:focal as dist
#COPY --from=build /usr/local/go /usr/local/go
COPY --from=build /g/main /usr/local/bin/obs-trtc
COPY --from=build /g/tmpl /usr/local/bin/tmpl
WORKDIR /usr/local/bin
CMD ["./obs-trtc"]

