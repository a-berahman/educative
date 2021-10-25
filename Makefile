.PHONY: postgres createdb dropdb migrateup migratedown

install-dev: ## Get the installation dependencies
	@go get -v -d ./...
	@go get -u github.com/canthefason/go-watcher
	@go install github.com/canthefason/go-watcher/cmd/watcher
	@go get -u golang.org/x/lint/golint
	@go get -u github.com/swaggo/swag/cmd/swag
postgres:
	docker run --name postgres-0 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -d -p 5432:5432 postgres:alpine	
createdb:
	docker exec -it postgres-0 createdb --username=postgres --owner=postgres test
dropdb:
	docker exec -it postgres-0 dropdb test
migration:
	go run main.go migrate -m migrations -c env.yaml
run:
	go run main.go serve -c env.yaml
migrateup:
	migrate -path migrations -database "postgresql://postgres:password@localhost:5432/test?sslmode=disable" -verbose up							   
migratedown:
	migrate -path migrations -database "postgresql://postgres:password@localhost:5432/test?sslmode=disable" -verbose down
test:
	go test ./...