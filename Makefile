postgres:
	docker run -dp 8082:5432 -e POSTGRES_PASSWORD=root -e POSTGRES_USER=root --name infosgroup-employee-management-db postgres:13-alpine
	docker exec -it infosgroup-employee-management-db apk add -U tzdata
	docker exec -it infosgroup-employee-management-db cp /usr/share/zoneinfo/America/Bogota /etc/localtime

createdb:
	docker exec -it infosgroup-employee-management-db createdb --username=root --owner=root infosgroup-employee-management-db

dropdb:
	docker exec -it infosgroup-employee-management-db dropdb infosgroup-employee-management-db

migrate:
	migrate create -ext sql -dir internal/database/migrations -seq init_schema

migrateup:
	migrate -path internal/database/migrations -database "postgresql://postgres:ZmK3tCJ6QxDI8lyjSn2g@employee-management.cddntvxjxrel.us-east-1.rds.amazonaws.com:5432/employeemanagement_db?sslmode=disable" -verbose up

migratedown:
	migrate -path internal/database/migrations -database "postgresql://postgres:ZmK3tCJ6QxDI8lyjSn2g@employee-management.cddntvxjxrel.us-east-1.rds.amazonaws.com:5432/employeemanagement_db?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./internal/database/tests -count=1

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test