package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"library-management-api/auth-service/gateway/grpc"
	"library-management-api/auth-service/init/database"
	"os"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	database.RunDB()
}

func main() {
	grpc.RunGRPC()
}
