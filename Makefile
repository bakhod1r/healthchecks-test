EMAIL ?= phapp0224mb@gmail.com

env: ## create .env from .env.example with a random secret
	@test -f .env || sed "s/^HC_SECRET_KEY=.*/HC_SECRET_KEY=$$(openssl rand -hex 32)/" .env.example > .env

.PHONY: env up down restart logs superuser ping lint

up: ## start healthchecks (:7077) and backend (:7071)
	docker compose up -d --force-recreate

down:
	docker compose down

restart: down up

logs:
	docker compose logs -f healthchecks

superuser: ## make superuser EMAIL=you@example.com
	docker compose exec -it healthchecks /opt/healthchecks/manage.py createsuperuser --email $(EMAIL)

ping: ## uses HC_UUID from .env
	./healthcheck.sh

lint:
	go vet ./...
	shellcheck healthcheck.sh
	docker compose config -q
