migrate:
	docker exec -i go_prod_mysql mysql -u root --password="test_pass" test < ./test.sql

migrate_prod:
	./migrateProd

migrate_test:
	./migrateTest

run_tests:
	docker exec -i awesomeproject_main_1 go test ./...

run_monkey_migration:
	docker exec -i awesomeproject_main_1 go run cmd/migrateDataFromMonkeyApi.go
