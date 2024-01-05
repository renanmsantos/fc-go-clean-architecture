createmigration:
	migrate create -ext=sql -dir=sql/migrations -seq init
migrate:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose up
migratedown:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose down

.PHONY: createmigration migrate migratedown