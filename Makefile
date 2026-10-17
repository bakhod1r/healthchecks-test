EMAIL ?= phapp0224mb@gmail.com

.PHONY: up down restart logs superuser ping lint

up: ## start healthchecks (:7077) and backend (:8080)
	docker compose up -d --force-recreate

down:
	docker compose down

restart: down up

logs:
	docker compose logs -f healthchecks

superuser: ## make superuser EMAIL=you@example.com
	docker compose exec -it healthchecks /opt/healthchecks/manage.py createsuperuser --email $(EMAIL)

ping: ## make ping HC_PING_URL=http://localhost:7077/ping/<uuid>
	./healthcheck.sh

lint:
	go vet ./...
	shellcheck healthcheck.sh
	docker compose config -q
