####
# This stage installs all the dependencies, builds the core binary.
####
FROM golang:1.21-alpine AS builder

# install the required dependencies: bash, git and make
RUN apk update \
    && apk add --no-cache bash git make

ENV USER=appuser
ENV UID=10001

# create appuser
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

# copy the source files.
WORKDIR /usr/app
COPY ../../../ .

# download and verify dependencies
RUN go mod download
RUN go mod verify

# build the binary
RUN make build

####
# This simply runs the application.
####
FROM scratch

WORKDIR /usr/app

# import the user and group files from the builder
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# copy our static executable
COPY --from=builder /usr/app/.build/core /usr/app/core

# use an unprivileged user
USER appuser:appuser

EXPOSE ${PORT}

# execute the binary
ENTRYPOINT ["/usr/app/core"]
