# healthchecks-test

Minimal playground for testing [healthchecks.io](https://healthchecks.io).

- `main.go` — backend, `GET /health` returns `{"status":"ok"}`.
- `healthcheck.sh` — cron script: sends one ping to `$HC_PING_URL`.

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

Remove the cron entry and the check goes down after period + grace.

| Env | Default |
|---|---|
| `HC_PING_URL` | required |

macOS: cron cannot read `~/Desktop` without Full Disk Access for `/usr/sbin/cron`; keep the script elsewhere or grant access.
