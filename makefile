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

proto-users:
	@protoc \
		--proto_path=users-service/api/pb "users-service/api/pb/user.proto" \
		--go_out=auth-service/pkg/proto --go_opt=paths=source_relative \
		--go_out=users-service/pkg/proto --go_opt=paths=source_relative \
		--go-grpc_out=auth-service/pkg/proto --go-grpc_opt=paths=source_relative \
		--go-grpc_out=users-service/pkg/proto --go-grpc_opt=paths=source_relative

proto-auth:
	@protoc \
		--proto_path=auth-service/api/pb "auth-service/api/pb/auth.proto" \
		--go_out=auth-service/pkg/proto --go_opt=paths=source_relative \
		--go_out=users-service/pkg/proto --go_opt=paths=source_relative \
		--go_out=books-service/pkg/proto --go_opt=paths=source_relative \
		--go-grpc_out=auth-service/pkg/proto --go-grpc_opt=paths=source_relative \
		--go-grpc_out=users-service/pkg/proto --go-grpc_opt=paths=source_relative \
		--go-grpc_out=books-service/pkg/proto --go-grpc_opt=paths=source_relative

.PHONY: docker-compose-db goose goose-create proto-books proto-users proto-auth
