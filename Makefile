export MYSQL_URL='mysql://root:superSecretPassword@tcp(localhost:3308)/db_forum'

migrate-create:
	@ migrate create -ext sql -dir scripts/migrations -seq $(name)

migrate-up:
	@ migrate -database ${MYSQL_URL} -path scripts/migrations -verbose up

migrate-down:
	@ migrate -database ${MYSQL_URL} -path scripts/migrations -verbose down
