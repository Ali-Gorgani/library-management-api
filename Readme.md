To run manually, run the following command in the terminal:
1. ```docker compose up``` for running databases.
2. ```go run api-gateway/cmd/main.go``` for running api-gateway server.
3. ```go run auth-service/cmd/main.go``` for running auth-service gRPC server.
4. ```go run users-service/cmd/main.go``` for running user-service gRPC server.
5. then go to http://localhost:8080/swagger for api documentation and test the APIs.