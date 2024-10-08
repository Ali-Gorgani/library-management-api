package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"library-management-api/users-service/configs"
	"library-management-api/users-service/gateway/grpc"
	"library-management-api/users-service/init/database"
	"os"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	configs.RunConfig("users-service")
	database.RunDB()
}

func main() {
	grpc.RunGRPC()
}
