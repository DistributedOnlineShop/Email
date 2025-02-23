package api

import (
	"github.com/go-redis/redis"

	"Email/pb"
	"Email/util"
)

type Server struct {
	pb.UnimplementedEmailServer
	config util.Config
	redis  *redis.Client
}

func ServerSetUp(config util.Config, rdb *redis.Client) (*Server, error) {
	return &Server{
		config: config,
		redis:  rdb,
	}, nil
}
