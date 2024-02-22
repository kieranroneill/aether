#!/usr/bin/env bash

SCRIPT_DIR=$(dirname "${0}")

source "${SCRIPT_DIR}"/set_vars.sh

# Public: Injects the version and runs the program.
#
# $1 - [optional] a version to inject, otherwise the version from the VERSION file is read.
#
# Examples
#
#   ./bin/run.sh # reads the version in the VERSION file
#   ./bin/run.sh "1.2.3"
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

  printf "%b running %b...\n" "${INFO_PREFIX}" "${APPLICATION_NAME}"
  go run -ldflags "-X main.Version=$version" cmd/"${APPLICATION_NAME}"/main.go


  printf "%b done!\n" "${INFO_PREFIX}"
}

# and so, it begins...
main "$1"
