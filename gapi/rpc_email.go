package gapi

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pbs "Email/pb/grpc_server"
)

func (s *Server) SendEmail(ctx context.Context, req *pbs.SignUpRequest) (*pbs.SignUpResponse, error) {
	data, err := s.EmailContextTpye(req.GetSituation(), req.GetEmail())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to Generate Email Context")
	}
	err = s.EmailHeader(data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Fail to Send Email")
	}

	return &pbs.SignUpResponse{Message: "Successfully sent the verification code email"}, nil
}

func (s *Server) VerifyEmail(ctx context.Context, req *pbs.VerifyEmailRequest) (*pbs.VerifyEmailResponse, error) {
	verificationCode, err := s.GetData(req.GetEmail())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Verification code not found")
	}

	if req.GetVerificationCode() == verificationCode {
		return &pbs.VerifyEmailResponse{Message: "Verification code correct"}, nil
	}

	return nil, status.Errorf(codes.Unauthenticated, "Verification code incorrect")
}
