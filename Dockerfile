FROM alpine:3.1

# update
RUN apk update

# need them to connect to AWS
RUN apk add curl ca-certificates && \
    update-ca-certificates

RUN apk add git

ADD app /
ADD settings.yml /

# temporary
CMD ["/app", "a"]
