#!/bin/bash

docker run \
  --rm \
  --name raspberrypiLocal-lightController \
  -p 8080:8080 \
  -v "$(pwd)"/src:/data \
  --workdir /data \
  --env-file=.env.prod \
  --env-file=.env.secrets \
  --net mqtt-network \
  -itd \
  raspberrypiLocal-lightController-img \
  /bin/bash -c "sh run.sh"
