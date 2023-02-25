FROM --platform=$BUILDPLATFORM golang:1.20 AS builder
ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETARCH
RUN echo "I am running on $BUILDPLATFORM, building for platform $TARGETPLATFORM, target arch $TARGETARCH"

WORKDIR /data
COPY ./src /data
RUN sh go-init.sh
RUN GOOS=linux GOARCH=$TARGETARCH go build -o light-controller

FROM --platform=$BUILDPLATFORM ubuntu:kinetic
ARG TARGETPLATFORM
ARG BUILDPLATFORM
RUN echo "I am running on $BUILDPLATFORM, building for $TARGETPLATFORM"
RUN apt-get update -y && apt-get upgrade -y && apt-get install -yq --no-install-recommends \
  locales \
  curl \
  tzdata \
  ca-certificates \
  openssl \
  nano
RUN apt-get clean && rm -rf /var/lib/apt/lists/*
RUN update-ca-certificates

RUN echo "en_US.UTF-8 UTF-8" > /etc/locale.gen && locale-gen
ENV LANG en_US.utf8
ENV TZ=Europe/Warsaw
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /root/
COPY --from=builder /data/light-controller ./
CMD ["./light-controller"]
