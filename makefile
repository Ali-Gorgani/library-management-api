docker-compose-db:
	docker compose exec -it db psql -U root -d library_db

goose:
	goose -dir migrations postgres "host=localhost port=5432 user=root password=secret dbname=library_db sslmode=disable" $(GOOSE_COMMAND)

GOOSE_COMMAND := $(if $(filter-out goose,$(MAKECMDGOALS)),$(filter-out goose,$(MAKECMDGOALS)),)

goose-create:
	goose -s -dir migrations create $(NAME) $(TYPE)

NAME := $(if $(filter-out goose-create,$(MAKECMDGOALS)),$(filter-out goose-create,$(MAKECMDGOALS)),)
TYPE := sql

%:
	@true

proto-books:
	rm -f books-service/api/pb/*.go
	protoc --proto_path=books-service/api/pb --go_out=books-service/api/pb --go_opt=paths=source_relative \
    --go-grpc_out=books-service/api/pb --go-grpc_opt=paths=source_relative \
    books-service/api/pb/api.proto

.PHONY: docker-compose-db goose goose-create proto-books
