package gapi

import (
	"github.com/go-redis/redis"

	pbe "Email/pb/email"
	pbs "Email/pb/session"
	pbv "Email/pb/verification"
	"Email/util"
)

type Server struct {
	pbe.UnimplementedEmailServer
	pbv.UnimplementedVerificationServer
	pbs.UnimplementedSessionServer

	config util.Config
	redis  *redis.Client
}

func ServerSetUp(config util.Config, rdb *redis.Client) (*Server, error) {
	return &Server{
		config: config,
		redis:  rdb,
	}, nil
}
