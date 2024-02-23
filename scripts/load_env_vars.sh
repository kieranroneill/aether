#!/usr/bin/env bash

SCRIPT_DIR=$(dirname "${0}")

source "${SCRIPT_DIR}"/set_vars.sh

# Public: Loads the .env file into environment variables.
#
# $1 - the name of the app's config file.
#
# Examples
#
#   ./scripts/load_env_vars.sh "core"
#
# Returns exit code 1 if no .env file exists, otherwise, exit code 0 is returned.
function main() {
  local env_file

  set_vars

  env_file=".env.${1}"

  # check if the config file exists
  if [[ ! -f "${CONFIG_DIR}/${env_file}" ]];
    then
      printf "%b no .env file exists for app %b at %b, is the app correct? \n" "${ERROR_PREFIX}" "${1}" "${CONFIG_DIR}/${env_file}"
      exit 1
  fi

  set -a
  # shellcheck source=./.config/.env.*
  source "${CONFIG_DIR}/${env_file}"
  set +a

  exit 0
}

# and so, it begins...
main "${1}"
