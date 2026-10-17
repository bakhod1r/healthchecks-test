#!/usr/bin/env bash
# Checks backend /health and reports the result to healthchecks.io.
#
# crontab -e:
#   * * * * * HC_PING_URL=https://hc-ping.com/<uuid> /path/to/healthcheck.sh >>/tmp/healthcheck.log 2>&1
#
# Env:
#   HC_PING_URL  required, check ping URL (hosted or self-hosted .../ping/<uuid>)
#   BACKEND_URL  default http://localhost:8080
set -uo pipefail

HC_PING_URL="${HC_PING_URL:?set HC_PING_URL}"
HC_PING_URL="${HC_PING_URL%/}"
BACKEND_URL="${BACKEND_URL:-http://localhost:8080}"

ping_hc() { # <suffix> [body]
  curl -fsS -m 10 --retry 3 -o /dev/null --data-raw "${2:-}" "$HC_PING_URL$1" \
    || echo "$(date '+%F %T') ping $1 failed" >&2
}

ping_hc /start

output="$(curl -fsS -m 5 "$BACKEND_URL/health" 2>&1)"
rc=$?

echo "$(date '+%F %T') rc=$rc $output"
ping_hc "/$rc" "$(printf '%s' "$output" | tail -c 10000)"
exit "$rc"
