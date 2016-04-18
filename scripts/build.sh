#!/bin/bash
set -x #echo on

rm app
GOOS=linux go build pr_helper/cmd/app
docker build -t pr_helper .
