package gapi

import (
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) SetData(key, contect string, expiration time.Duration) error {
	cmd := s.redis.Set(key, contect, expiration)
	if err := cmd.Err(); err != nil {
		return status.Errorf(codes.Internal, "Failed to save to Redis")
	}
	return nil
}

func (s *Server) GetData(key string) (string, error) {
	cmd := s.redis.Get(key)
	if err := cmd.Err(); err != nil {
		return "", status.Errorf(codes.Internal, "Failed to taking data in Redis")
	}
	return cmd.Val(), nil
}

func (s *Server) DeleteData(key string) error {
	cmd := s.redis.Del(key)
	if err := cmd.Err(); err != nil {
		return status.Errorf(codes.Internal, "Failed to delete data in Redis")
	}

	return nil
}
