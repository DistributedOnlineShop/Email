package gapi

import (
	"github.com/go-redis/redis"

	pbs "Email/pb/grpc_server"
	"Email/util"
)

type Server struct {
	pbs.UnimplementedEmailServer
	config util.Config
	redis  *redis.Client
}

func ServerSetUp(config util.Config, rdb *redis.Client) (*Server, error) {
	return &Server{
		config: config,
		redis:  rdb,
	}, nil
}
