ARG STEP_1_IMAGE=golang:1.14.1-alpine3.11
ARG STEP_2_IMAGE=alpine:3.11

# -----------------------------------------------------------------------------
# STEP_1
# -----------------------------------------------------------------------------
FROM ${STEP_1_IMAGE} AS STEP_1

ARG GIT_BRANCH="master"

# Add github
RUN apk update && \
    apk upgrade && \
    apk add --no-cache \
        git \
        make \
        bash

COPY . /go/src/azure-file-fetch

# Checkout to target branch and if not master, then locally test
# merging master into branch.
WORKDIR /go/src/azure-file-fetch

RUN make build

# -----------------------------------------------------------------------------
# STEP_2
# -----------------------------------------------------------------------------
FROM ${STEP_2_IMAGE} AS STEP_2

COPY --from=STEP_1 /go/bin/azure-file-fetch /usr/bin

# Create Terraform User
RUN addgroup -S azure-file-fetch && adduser -S azure-file-fetch -G azure-file-fetch

USER azure-file-fetch

WORKDIR /home/azure-file-fetch

ENTRYPOINT ["azure-file-fetch"]