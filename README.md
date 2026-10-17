# healthchecks-test

Minimal playground for testing [healthchecks.io](https://healthchecks.io).

- `main.go` — backend, `GET /health` returns `{"status":"ok"}`.
- `healthcheck.sh` — cron script: pings `/start`, checks `$BACKEND_URL/health`, then pings `/<exit-code>` (`/0` = success) with the output as body.

## Run

```sh
docker compose up -d   # healthchecks on :8000, backend on :8080
```

Open http://localhost:8000, sign up, add a check (period 1 min, grace 1 min), copy its ping URL.

Test once by hand:

```sh
HC_PING_URL=http://localhost:8000/ping/<uuid> ./healthcheck.sh
```

Then `crontab -e`:

```cron
* * * * * HC_PING_URL=http://localhost:8000/ping/<uuid> /path/to/healthcheck.sh >>/tmp/healthcheck.log 2>&1
```

Stop the backend (`docker compose stop backend`) and the check goes down.

| Env | Default |
|---|---|
| `HC_PING_URL` | required |
| `BACKEND_URL` | `http://localhost:8080` |

macOS: cron cannot read `~/Desktop` without Full Disk Access for `/usr/sbin/cron`; keep the script elsewhere or grant access.
