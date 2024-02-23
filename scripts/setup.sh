#!/usr/bin/env bash

SCRIPT_DIR=$(dirname "${0}")

source "${SCRIPT_DIR}"/set_vars.sh

# Public: Creates a .env.* file for an application, if it don't exist.
#
# $1 - the name of the web.
#
# Examples
#
#   ./scripts/setup.sh "core"
#
# Returns exit code 1 if no example file for the web exists, otherwise, exit code 0 is returned.
function main() {
  local env_example_file
  local env_file

  set_vars

  env_example_file=".env.${1}.example"

  if [[ ! -f "${CONFIGS_DIR}/${env_example_file}" ]];
    then
      printf "%b no example at %b, is the application correct? \n" "${ERROR_PREFIX}" "${env_example_file}"
      exit 1
  fi

  # create a config directory if one doesn't exist
  if [[ ! -d "${CONFIG_DIR}" ]];
    then
      printf "%b creating new %b directory... \n" "${INFO_PREFIX}" "${CONFIG_DIR}"
      mkdir -p "${CONFIG_DIR}"
  fi

  env_file=".env.${1}"

  printf "%b creating %b files...\n" "${INFO_PREFIX}" "${env_file}"

  # create the .env.* file
  cp -n "${CONFIGS_DIR}/${env_example_file}" "${CONFIG_DIR}/${env_file}"

  printf "%b done!\n" "${INFO_PREFIX}"

  exit 0
}

# and so, it begins...
main "${1}"
