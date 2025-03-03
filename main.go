package main

import (
	"net"
	"os"

	"github.com/go-redis/redis"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"Email/gapi"
	pbe "Email/pb/email"
	pbv "Email/pb/verification"
	"Email/util"
)

func main() {
	config, err := util.LoadConfig("./")
	if err != nil {
		log.Error().Err(err).Msg("app.env is not found")
		os.Exit(1)
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	rdb := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddress,
		Password: "",
		DB:       1,
	})

	server, err := gapi.ServerSetUp(config, rdb)
	if err != nil {
		log.Error().Err(err).Msg("failed to set up server")
		os.Exit(1)
	}

	Logger := grpc.UnaryInterceptor(gapi.GrpcLogger)
	grpcServer := grpc.NewServer(Logger)

	pbe.RegisterEmailServer(grpcServer, server)
	pbv.RegisterVerificationServer(grpcServer, server)

	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.EmailPort)
	if err != nil {
		log.Error().Err(err).Msg("failed to listen")
		os.Exit(1)
	}

	log.Info().Msgf("Connect to grpc server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Error().Err(err).Msg("failed to connecting serve")
		os.Exit(1)
	}
}
