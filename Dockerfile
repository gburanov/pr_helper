FROM alpine:3.1

# need them to connect to AWS
RUN apk --update upgrade && \
    apk add curl ca-certificates && \
    update-ca-certificates && \
    rm -rf /var/cache/apk/*

ADD app /
CMD ["/app"]
