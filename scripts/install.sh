#!/usr/bin/env bash

SCRIPT_DIR=$(dirname "${0}")

source "${SCRIPT_DIR}"/set_vars.sh

# Public: Creates env files if they don't exist.
#
# Examples
#
#   ./scripts/install.sh
#
# Returns exit code 0.
function main() {
  set_vars

  printf "%b creating env files...\n" "${INFO_PREFIX}"

  # create the .env.* files
  cp -n "${CONFIGS_DIR}/.env.core.example" "${CONFIGS_DIR}/.env.core"

  printf "%b done!\n" "${INFO_PREFIX}"
}

# and so, it begins...
main
