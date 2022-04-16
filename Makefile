prepare_project:
	cp .env.example .env
	cp .env.testing.example .env.testing
	docker-compose build
	docker-compose up -d
	ENV=production ./scripts/migrate
	ENV=testing ./scripts/migrate

migrate_prod:
	ENV=production ./scripts/migrate

migrate_test:
	ENV=testing ./scripts/migrate

run_tests:
	docker exec -i awesomeproject_main_1 go test ./...

run_monkey_migration:
	docker exec -i awesomeproject_main_1 go run cmd/migrateDataFromMonkeyApi.go
