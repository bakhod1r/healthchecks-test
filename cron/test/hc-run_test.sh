#!/usr/bin/env bash
# Tests hc-run.sh with a stub curl that logs its arguments.
set -uo pipefail

here="$(cd "$(dirname "$0")" && pwd)"
run="$here/../hc-run.sh"
tmp="$(mktemp -d)"
trap 'rm -rf "$tmp"' EXIT

cat >"$tmp/curl" <<'EOF'
#!/usr/bin/env bash
last="${!#}"
body=""
for a in "$@"; do [ "$a" = "@-" ] && body="$(cat)"; done
echo "$last|$body" >>"$CURL_LOG"
EOF
chmod +x "$tmp/curl"
export PATH="$tmp:$PATH" CURL_LOG="$tmp/log"

fails=0
check() { # name expected actual
  if [ "$2" = "$3" ]; then echo "ok   $1"; else echo "FAIL $1: want [$2] got [$3]"; fails=$((fails + 1)); fi
}

: >"$CURL_LOG"
"$run" http://hc/ping/abc/ echo hello >/dev/null
check "success exit code" 0 "$?"
check "success pings" $'http://hc/ping/abc/start|\nhttp://hc/ping/abc/0|hello' "$(cat "$CURL_LOG")"

: >"$CURL_LOG"
"$run" http://hc/ping/abc sh -c 'echo boom >&2; exit 3' >/dev/null
check "fail exit code" 3 "$?"
check "fail pings" $'http://hc/ping/abc/start|\nhttp://hc/ping/abc/3|boom' "$(cat "$CURL_LOG")"

"$run" only-url >/dev/null 2>&1
check "usage error" 64 "$?"

if [ "$fails" -ne 0 ]; then
  echo "$fails failed"
  exit 1
fi
echo PASS
