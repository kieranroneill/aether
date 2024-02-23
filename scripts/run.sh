#!/usr/bin/env bash

SCRIPT_DIR=$(dirname "${0}")

source "${SCRIPT_DIR}"/set_vars.sh

# Public: Checks the apps have been configured correctly, then starts Docker Compose.
#
# Examples
#
#   ./scripts/run.sh
#
# Returns exit code 1 if an app has not been configured correctly, otherwise, exit code 0 is returned.
function main() {
  set_vars

  # check core configuration
  if [ ! -f .config/.env.core ]; then
    printf "\n%b core application not configured correctly, have you run 'make install'?" "${ERROR_PREFIX}"

    exit 1
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
