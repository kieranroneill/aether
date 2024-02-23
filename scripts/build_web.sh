#!/usr/bin/env bash

SCRIPT_DIR=$(dirname "${0}")

source "${SCRIPT_DIR}"/set_vars.sh

# Public: Builds the the Next.js app to the .next/ directory.
#
# Examples
#
#   ./scripts/build_web.sh
#
# Returns exit code 0.
function main() {
  printf "%b building web app...\n" "${INFO_PREFIX}"
  yarn build

  printf "%b done!\n" "${INFO_PREFIX}"

  exit 0
}

# and so, it begins...
main "$1"
