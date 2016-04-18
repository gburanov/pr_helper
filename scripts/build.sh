#!/bin/bash
set -x #echo on

rm web
GOOS=linux go build pr_helper/cmd/web
docker build -t pr_helper .
