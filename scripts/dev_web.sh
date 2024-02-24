#!/usr/bin/env bash

SCRIPT_DIR=$(dirname "${0}")

source "${SCRIPT_DIR}"/set_vars.sh

# Public: Runs Next.js in development mode.
#
# Examples
#
#   ./bin/dev_web.sh
#
# Returns exit code 0.
function main() {
  set_vars

  printf "%b starting web server...\n" "${INFO_PREFIX}"
  yarn dev

  printf "%b done!\n" "${INFO_PREFIX}"

  exit 0
}

# and so, it begins...
main
