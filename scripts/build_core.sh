#!/usr/bin/env bash

SCRIPT_DIR=$(dirname "${0}")

source "${SCRIPT_DIR}"/set_vars.sh

# Public: Inject the version and builds the binary to .build/core.
#
# $1 - [optional] a version to inject, otherwise the version from the VERSION file is read.
#
# Examples
#
#   ./scripts/build_core.sh # reads the version in the VERSION file
#   ./scripts/build_core.sh "1.2.3"
#
# Returns exit code 0.
function main() {
  local version

  set_vars

  version=$(<VERSION)

  # if the version argument exists, use it
  if [ -n "$1" ]; then
    version="$1"
  fi

  printf "%b compiling core binary...\n" "${INFO_PREFIX}"
  go build -o "${BUILD_DIR}"/core -ldflags "-X main.Version=$version" "${CORE_SRC_DIR}"/main.go

  printf "%b done!\n" "${INFO_PREFIX}"

  exit 0
}

# and so, it begins...
main "$1"
