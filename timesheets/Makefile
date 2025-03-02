.PHONY: build

build:
	sam build --config-env env.local.json
	
deploy:
	sam build --config-env env.deploy.json
	sam deploy --stack-name timesheets-service --no-confirm-changeset --no-fail-on-empty-changeset

local: build up
	sam local start-api -n env.local.json --warm-containers eager --docker-network lambda-local

up:
	docker-compose up -d
down:
	docker-compose down
	
sync: build
	sam sync
watch: build
	sam sync --watch
