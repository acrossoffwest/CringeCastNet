#!/bin/bash

docker build --platform=linux/amd64 -t cringecast-client-linux-build .
docker run --rm -it --platform=linux/amd64 -v $(pwd)/../bin:/output cringecast-client-linux-build /bin/sh -c "cp /app/cringecast-client-linux /output"

docker build --platform=linux/arm64 -t cringecast-client-linux-build-arm .
docker run --rm -it --platform=linux/arm64 -v $(pwd)/../bin:/output cringecast-client-linux-build-arm /bin/sh -c "cp /app/cringecast-client-linux /output/cringecast-client-linux-arm"
