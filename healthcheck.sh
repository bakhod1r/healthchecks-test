#!/usr/bin/env bash
# Sends a ping to healthchecks.
# Reads .env next to the script; HC_PING_URL overrides, else built from HC_UUID + HC_PORT.
#
# crontab -e:
#   * * * * * /path/to/healthcheck.sh >>/tmp/healthcheck.log 2>&1
set -euo pipefail

env_file="$(dirname "$0")/.env"
if [ -f "$env_file" ]; then
  set -a
  # shellcheck source=/dev/null
  . "$env_file"
  set +a
fi

HC_PING_URL="${HC_PING_URL:-http://localhost:${HC_PORT:-7077}/ping/${HC_UUID:?set HC_UUID in .env}}"

curl -fsS -m 10 --retry 3 "$HC_PING_URL"
echo " $(date '+%F %T') pinged"
