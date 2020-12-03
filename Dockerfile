FROM golang:1-alpine AS builder
MAINTAINER Luc Michalski <lmichalski@evolutive-business.com>

ENV GOOS=linux \
    GO111MODULE=on

WORKDIR /go/src/github.com/percona/pmm-admin

COPY . /go/src/github.com/percona/pmm-admin

# gcc/g++ are required to build SASS libraries for extended version
RUN apk update && \
    apk add --no-cache git ca-certificates make

RUN make release-gomod

FROM alpine:3.12 AS runtime 

# Build arguments
ARG TINI_VERSION=${TINI_VERSION:-"v0.19.0"}

# Install tini to /usr/local/sbin
ADD https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini-muslc-amd64 /usr/local/sbin/tini

# Install runtime dependencies & create runtime user
RUN apk --no-cache --no-progress add ca-certificates \
    && chmod +x /usr/local/sbin/tini \
    && mkdir -p /opt \
    && adduser -D percona -h /opt/pmm -s /bin/sh \
    && su percona -c 'cd /opt/pmm; mkdir -p bin data config'

# Switch to user context
USER percona
WORKDIR /opt/pmm

# Copy pmm-admin binary to /opt/pmm/bin
COPY --from=builder /go/src/github.com/percona/pmm-admin/bin/pmm-admin /opt/pmm/bin/pmm-admin
ENV PATH $PATH:/opt/pmm/bin

# Container configuration
VOLUME ["/opt/pmm/data", "/opt/pmm/config"]
ENTRYPOINT ["tini", "-g", "--", "pmm-admin"]
