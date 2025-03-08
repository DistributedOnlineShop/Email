package gapi

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	pbs "Email/pb/session"
)

func (s *Server) CreationSessionId(ctx context.Context, req *pbs.CreateSessionIdRequest) (*pbs.CreateSessionIdResponse, error) {
	conn, err := grpc.NewClient(s.config.UsersManagementPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to connect Email Server")
	}
	defer conn.Close()

	client := pbs.NewSessionClient(conn)

	sessionRes, err := client.CreateSessionId(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error during verification")
	}

	return sessionRes, nil
}
