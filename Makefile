postgresconsole:
	docker exec -it postgres17 psql -U root -d bp_dev
postgresrun:
	docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=postgres -d postgres:17-alpine

postgresstart:
	docker start postgres17

postgresstop:
	docker stop postgres17

createdb:
	docker exec -it postgres17 createdb --username=root --owner=root bp_dev

dropdb:
	docker exec -it postgres17 dropdb bp_dev


newmigration:
	migrate create -ext sql -dir db/migration -seq $(name)

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up
migrateupone:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down
migratedownone:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

sqlc:
	sqlc generate

server:
	go run main.go


.PHONY: postgresconsole postgresrun postgresstart postgresstop createdb dropdb newmigration migrateup migrateupone migratedown migratedownone server sqlc