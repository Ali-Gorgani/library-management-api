To run manually, run the following command in the terminal:
1. ```docker compose up``` for running databases.
2. ```go run util/cli/main.go api-gateway``` for running api-gateway server.
3. ```go run util/cli/main.go auth-service``` for running auth-service gRPC server.
4. ```go run util/cli/main.go users-service``` for running users-service gRPC server.
5. then go to http://localhost:8080/swagger for api documentation and test the APIs.