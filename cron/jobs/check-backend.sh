#!/usr/bin/env bash
# Fails when the backend /health endpoint is not 200.
set -euo pipefail
curl -fsS -m 5 "${BACKEND_URL:-http://localhost:8080}/health"
