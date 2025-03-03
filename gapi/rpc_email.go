package gapi

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pbe "Email/pb/email"
)

func (s *Server) SendEmail(ctx context.Context, req *pbe.SendEmailRequest) (*pbe.SendEmailResponse, error) {
	data, err := s.EmailContextTpye(req.GetSituation(), req.GetEmail())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to Generate Email Context")
	}
	err = s.EmailHeader(data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Fail to Send Email")
	}

	return &pbe.SendEmailResponse{Message: "Successfully sent the verification code email"}, nil
}

func (s *Server) VerifyEmail(ctx context.Context, req *pbe.VerifyEmailRequest) (*pbe.VerifyEmailResponse, error) {
	verificationCode, err := s.GetData(req.GetEmail())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Verification code not found")
	}

	if req.GetVerificationCode() == verificationCode {
		return &pbe.VerifyEmailResponse{Message: "Verification code correct"}, nil
	}

	return nil, status.Errorf(codes.Unauthenticated, "Verification code incorrect")
}
