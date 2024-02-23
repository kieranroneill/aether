#!/usr/bin/env bash

SCRIPT_DIR=$(dirname "${0}")

source "${SCRIPT_DIR}"/set_vars.sh

# Public: Creates .env.* files for each application, if they don't exist and installs yarn dependencies.
#
# Examples
#
#   ./scripts/install.sh
#
# Returns exit code 0.
function main() {
  set_vars

  if [[ ! -d "${CONFIG_DIR}" ]];
    then
      printf "%b creating new %b directory... \n" "${INFO_PREFIX}" "${CONFIG_DIR}"
      mkdir -p "${CONFIG_DIR}"
  fi

  printf "%b creating env files...\n" "${INFO_PREFIX}"

  # create the .env.* files
  cp -n "${CONFIGS_DIR}/.env.core.example" "${CONFIG_DIR}/.env.core"

  printf "%b installing dependencies...\n" "${INFO_PREFIX}"

  # install dependencies
  yarn install

  printf "%b done!\n" "${INFO_PREFIX}"

  exit 0
}

# and so, it begins...
main
