#!/bin/bash
set -x #echo on

GOOS=linux go build pr_helper/cmd/app
docker build -t pr_helper .
