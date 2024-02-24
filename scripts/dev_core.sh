#!/usr/bin/env bash

SCRIPT_DIR=$(dirname "${0}")

source "${SCRIPT_DIR}"/set_vars.sh

# Public: Injects the version and runs the Go core server.
#
# $1 - [optional] a version to inject, otherwise the version from the VERSION file is read.
#
# Examples
#
#   ./bin/dev_core.sh # reads the version in the VERSION file
#   ./bin/dev_core.sh "1.2.3"
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

  VERSION=$version

  export VERSION

  printf "%b starting core server...\n" "${INFO_PREFIX}"
  CompileDaemon -build="go build -o ${BUILD_DIR}/core ${CORE_SRC_DIR}/main.go" -command="${BUILD_DIR}/core"

  printf "%b done!\n" "${INFO_PREFIX}"

  exit 0
}

# and so, it begins...
main "$1"
