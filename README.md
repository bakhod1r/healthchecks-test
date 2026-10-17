# healthchecks-test

Minimal playground for testing [healthchecks.io](https://healthchecks.io).

- `main.go` — backend, `GET /health` returns `{"status":"ok"}`; with `HC_PING_URL` set it also pings healthchecks every `HC_PING_INTERVAL` (default `1m`). In compose use `HC_PING_URL=http://healthchecks:8000/ping/<uuid>`.
- `healthcheck.sh` — cron script: sends one ping to `$HC_PING_URL`.

## Run

```sh
cp .env.example .env   # set HC_SECRET_KEY, HC_UUID
docker compose up -d   # healthchecks on :7077, backend on :7071
```

Open http://localhost:7077, sign up, add a check (period 1 min, grace 1 min), copy its ping URL.

Test once by hand:

```sh
./healthcheck.sh   # reads HC_UUID from .env
```

Then `crontab -e`:

```cron
* * * * * /path/to/healthcheck.sh >>/tmp/healthcheck.log 2>&1
```

Remove the cron entry and the check goes down after period + grace.

| Env | Default |
|---|---|
| `HC_PING_URL` | required |

macOS: cron cannot read `~/Desktop` without Full Disk Access for `/usr/sbin/cron`; keep the script elsewhere or grant access.
