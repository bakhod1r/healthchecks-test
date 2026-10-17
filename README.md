# healthchecks-test

Minimal playground for testing [healthchecks.io](https://healthchecks.io): a tiny Go backend and a bash cronjob that reports every run to Healthchecks.

- `main.go` — backend, `GET /health` returns `{"status":"ok"}`.
- `cron/hc-run.sh <ping-url> <command...>` — wrapper: pings `/start`, runs the command, then pings `/<exit-code>` (`/0` = success) with the last 10 KB of output as the body.
- `cron/jobs/check-backend.sh` — the job: `curl -f $BACKEND_URL/health`.
- `cron/crontab` — runs the job every minute (alpine `crond` in the `cron` container).

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

The check turns green. Run `docker compose stop backend`: the next run pings `/7` (curl connect error), so the check goes down.

Without Docker, add to your own crontab:

```cron
* * * * * HC_PING_URL=https://hc-ping.com/<uuid> /path/to/cron/hc-run.sh "$HC_PING_URL" /path/to/cron/jobs/check-backend.sh
```

## Test

```sh
go vet ./...
bash cron/test/hc-run_test.sh
```
