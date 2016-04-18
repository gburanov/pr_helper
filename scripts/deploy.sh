#!/bin/bash
set -x #echo on

TAG=$(git rev-parse HEAD)

BuildAndPush() {
  PARENT_DIR=$(basename "${PWD%/*}")
  CURRENT_DIR="${PWD##*/}"
  IMAGE_NAME="gburanov/$CURRENT_DIR"
  docker build -t ${IMAGE_NAME}:${TAG} .
  docker tag ${IMAGE_NAME}:${TAG} ${IMAGE_NAME}:latest
  docker push ${IMAGE_NAME}
}

rm web
GOOS=linux go build pr_helper/cmd/web

BuildAndPush
