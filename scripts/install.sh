#!/usr/bin/env bash

SCRIPT_DIR=$(dirname "${0}")

source "${SCRIPT_DIR}"/set_vars.sh

# Public: Installs yarn and golang dependencies.
#
# Examples
#
#   ./scripts/install.sh
#
# Returns exit code 0.
function main() {
  set_vars

  printf "%b installing yarn dependencies...\n" "${INFO_PREFIX}"

  # install yarn dependencies
  yarn install

  printf "%b installing golang dependencies...\n" "${INFO_PREFIX}"

  # install golang dependencies
  go mod download
  go mod verify

  printf "%b done!\n" "${INFO_PREFIX}"

  exit 0
}

# and so, it begins...
main
