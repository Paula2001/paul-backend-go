create_database:
	docker exec -it go_prod_mysql mysql -u root --password="test_pass" -e "create database test"

migrate:
	docker exec -i go_prod_mysql mysql -u root --password="test_pass" test < ./test.sql

create_migrate:
	./migrate

run_tests:
	docker exec -i awesomeproject_main_1 go test ./...