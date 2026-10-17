#!/usr/bin/env bash
# Usage: hc-run.sh <ping-url> <command> [args...]
# Sends /start, runs the command, then reports its exit code (/0 = success)
# with the last 10 KB of output as the ping body.
set -uo pipefail

if [ $# -lt 2 ]; then
  echo "usage: $0 <ping-url> <command> [args...]" >&2
  exit 64
fi

url="${1%/}"
shift

ping() {
  curl -fsS -m 10 --retry 3 -o /dev/null "$@" || echo "hc-run: ping failed: $*" >&2
}

ping "$url/start"
output="$("$@" 2>&1)"
rc=$?
printf '%s\n' "$output"
printf '%s' "$output" | tail -c 10000 | ping --data-binary @- "$url/$rc"
exit "$rc"
