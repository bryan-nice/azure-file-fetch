ARG STEP_1_IMAGE=golang:1.14.1-alpine3.11
ARG STEP_2_IMAGE=alpine:3.11

# -----------------------------------------------------------------------------
# STEP_1
# -----------------------------------------------------------------------------
FROM ${STEP_1_IMAGE} AS STEP_1

ARG GITHUB_HEAD_REF="master"
ARG GITHUB_EVENT_NAME="push"

# Add github
RUN apk update && \
    apk upgrade && \
    apk add --no-cache \
        git \
        make \
        bash

WORKDIR /go/src

# Clone repo
RUN git clone https://github.com/bryan-nice/azure-file-fetch.git

# Checkout to target branch and if not master, then locally test
# merging master into branch.
WORKDIR /go/src/azure-file-fetch
RUN git checkout ${GITHUB_HEAD_REF}; \
    if [[ "${GITHUB_HEAD_REF}" != "master" && ${GITHUB_EVENT_NAME} == "pull_request" ]]; then \
        git merge master; \
    fi

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