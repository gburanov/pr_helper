# pr_helper

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

Copy blacklist file

And run it

    $ ./pr_helper a
