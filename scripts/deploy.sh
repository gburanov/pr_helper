#!/bin/bash
set -x #echo on


./scripts/build.sh

CURRENT_DIR="${PWD##*/}"
IMAGE_NAME="gburanov/$CURRENT_DIR"
docker push ${IMAGE_NAME}
