bash: ## Open shell
	docker compose exec go sh

run: ## Target to run main Go package
	docker compose exec go sh -c 'go run main.go'

build: ## Target to compile code into an executable
	docker compose exec go sh -c 'go build main.go'

run-exec: ## Target to run executable
	docker compose run go sh -c './main'