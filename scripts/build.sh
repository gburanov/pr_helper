#!/bin/bash
set -x #echo on

TAG=$(git rev-parse HEAD)

rm sqs
GOOS=linux go build pr_helper/cmd/sqs

PARENT_DIR=$(basename "${PWD%/*}")
CURRENT_DIR="${PWD##*/}"
IMAGE_NAME="gburanov/$CURRENT_DIR"
docker build -t ${IMAGE_NAME}:${TAG} .
docker tag ${IMAGE_NAME}:${TAG} ${IMAGE_NAME}:latest
