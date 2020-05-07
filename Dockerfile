#golang 1.14.2-alpine3.11
#recommeded always pulling image using digest
FROM golang@sha256:9b3ad7928626126b72b916609ad081cfb6c0149f6e60cef7fc1e9e15a0d1e973 AS builder

ARG REPO
ARG BRANCH

# Install git.
# Git is required for fetching the dependencies & sourcecode

RUN apk update && apk add --no-cache git bash

WORKDIR $GOPATH/

#tweak for disable cache
ARG CACHEBUST=1

RUN git clone $REPO -b $BRANCH --single-branch app

WORKDIR $GOPATH/app

# Using go get.
RUN go get -d -v

#For small binnary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags='-w -s -extldflags "-static"' -a -o /go/bin/main .

#FROM scratch
FROM alpine:latest
COPY --from=builder /go/bin/main /

# Pick One method whether to be able to access or not
# Execute the binary.
ENTRYPOINT ["/main"]

# Use to debug
#CMD ["/main"]
EXPOSE 1338
