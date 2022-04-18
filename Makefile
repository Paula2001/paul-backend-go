prepare_project:
	cp .env.example .env
	cp .env.testing.example .env.testing
	docker-compose build

start_project:
	docker-compose up -d

end_project:
	docker-compose down

migrate_prod:
	ENV=production ./scripts/migrate

migrate_test:
	ENV=testing ./scripts/migrate

run_tests:
	docker exec -it awesomeproject_main_1 bash -c "export ENV=testing && go test ./..."

run_monkey_migration:
	docker exec -i awesomeproject_main_1 go run cmd/migrateDataFromMonkeyApi.go
