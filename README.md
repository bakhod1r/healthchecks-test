# healthchecks-test

Minimal playground for testing [healthchecks.io](https://healthchecks.io): a tiny Go backend and one cronjob that reports every run to Healthchecks.

- `cmd/server` — backend, `GET /health` returns `{"status":"ok"}`.
- `cmd/cron` — every `CRON_INTERVAL` calls `BACKEND_URL/health`, sends `/start` before the run, then success or `/fail` (with error text) after.
- `internal/monitoring` — `domain` (check + signals), `application` (`RunMonitored`), `infrastructure` (HTTP pinger).

## Run locally

```sh
docker compose up -d healthchecks backend
```

1. Open http://localhost:8000, sign up, add a check (period 1 min, grace 1 min).
2. Copy its ping URL and swap the host for the container name: `http://healthchecks:8000/ping/<uuid>`.
3. Start the cron:

```sh
HC_PING_URL=http://healthchecks:8000/ping/<uuid> docker compose up -d cron
```

The check turns green. Run `docker compose stop backend` and the next run sends `/fail`, so the check goes down.

With hosted healthchecks.io, use `HC_PING_URL=https://hc-ping.com/<uuid>` and run `go run ./cmd/cron`.

## Env

| Var | Default |
|---|---|
| `HC_PING_URL` | required |
| `BACKEND_URL` | `http://localhost:8080` |
| `CRON_INTERVAL` | `1m` |
| `ADDR` (server) | `:8080` |

## Test

```sh
go test ./...
```
