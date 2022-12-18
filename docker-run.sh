#!/bin/bash

docker run \
  --rm \
  --name raspberrypiLocal-lightController \
  -v "$(pwd)"/src:/data \
  --workdir /data \
  --env-file=.env \
  --env-file=.env.dev \
  --env-file=.env.secrets \
  --net mqtt-network \
  -itd \
  c4stus/raspberrypi:lightcontroller \
  /bin/bash -c "sh run.sh"
