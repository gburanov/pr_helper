FROM alpine:3.1

# update
RUN apk update

# need them to connect to AWS
RUN apk add curl ca-certificates && \
    update-ca-certificates

RUN apk add git

ADD sqs /
ADD blacklist /
ADD whitelist /
#ADD cmd/web/index.gtpl cmd/web/index.gtpl
ADD settings.yml /

CMD ["/sqs"]
