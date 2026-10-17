#!/usr/bin/env bash
# Sends a ping to healthchecks.io.
#
# crontab -e:
#   * * * * * HC_PING_URL=https://hc-ping.com/<uuid> /path/to/healthcheck.sh >>/tmp/healthcheck.log 2>&1
set -euo pipefail

: "${HC_PING_URL:?set HC_PING_URL}"

curl -fsS -m 10 --retry 3 "$HC_PING_URL"
echo " $(date '+%F %T') pinged"
