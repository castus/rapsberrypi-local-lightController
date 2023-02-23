#!/bin/bash

docker run \
  --rm \
  --name raspberrypiLocal-lightController \
  --env-file=.env \
  --env-file=.env.dev \
  --env-file=.env.secrets \
  --net mqtt-network \
  -itd \
  c4stus/raspberrypi:lightcontroller
