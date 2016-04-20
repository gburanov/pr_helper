#!/bin/bash
set -x #echo on

rm sqs
GOOS=linux go build pr_helper/cmd/sqs
docker build -t pr_helper .
