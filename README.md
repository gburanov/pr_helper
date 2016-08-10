# pr_helper

[![Build Status](https://travis-ci.org/gburanov/pr_helper_frontend.svg?branch=master)](https://travis-ci.org/gburanov/pr_helper_frontend)

Installation and usage
----------------------
First of all, install Go

    brew install golang

Setup the $GOPATH variable

    mkdir ~/golang
    export GOPATH=~/golang
    export PATH=$PATH:~/golang/bin/

Download the project

    go get github.com/gburanov/pr_helper

Create settings.yml file

    authtoken: token
    repositorypath: path to repository

Set aws keys

   AWS_PRIVATE_ACCESS_KEY_ID
   AWS_PRIVATE_SECRET_ACCESS_KEY

Copy blacklist file

And run it

    $ pr_helper a
