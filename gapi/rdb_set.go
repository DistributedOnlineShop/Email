package gapi

import (
	"fmt"
	"time"
)

func (s *Server) SetData(key, contect string, expiration time.Duration) error {
	cmd := s.redis.Set(key, contect, expiration)
	if err := cmd.Err(); err != nil {
		return fmt.Errorf("Failed to save to Redis: %s", err)
	}
	return nil
}
