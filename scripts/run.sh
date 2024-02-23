#!/usr/bin/env bash

SCRIPT_DIR=$(dirname "${0}")

source "${SCRIPT_DIR}"/set_vars.sh

# Public: Checks the apps have been configured correctly, then starts Docker Compose.
#
# This script is used as the entry point for running the application in development mode and is not intended for
# production.
#
# Examples
#
#   ./scripts/run.sh
#
# Returns exit code 1 if an web has not been configured correctly, otherwise, exit code 0 is returned.
function main() {
  set_vars

  # check core configuration
  if [ ! -f "${CONFIG_DIR}"/.env.core ]; then
    printf "\n%b core application not configured correctly, have you run 'make setup'?" "${ERROR_PREFIX}"

    exit 1
  else
    source "${CONFIG_DIR}"/.env.core

    CORE_APP_PORT="${PORT}"

    export CORE_APP_PORT
  fi

  # check web configuration
  if [ ! -f "${CONFIG_DIR}"/.env.web ]; then
    printf "\n%b web application not configured correctly, have you run 'make setup'?" "${ERROR_PREFIX}"

    exit 1
  else
    source "${CONFIG_DIR}"/.env.web

    WEB_APP_PORT="${PORT}"

    export WEB_APP_PORT
  fi

  printf "\n%b starting docker compose...\n" "${INFO_PREFIX}"
  docker compose \
    -f ./deployments/docker-compose.yml \
    up \
    --build

  exit 0
}

# and so, it begins...
main
