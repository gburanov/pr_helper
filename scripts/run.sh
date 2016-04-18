#!/bin/bash
set -x #echo on

docker run \
  -p 80:3000 \
  -e GITHUB_ACCESS_TOKEN=$GITHUB_ACCESS_TOKEN \
  -v /repository \
  pr_helper:latest
