package gapi

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pbe "Email/pb/email"
	pbs "Email/pb/session"
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

	return &pbe.SendEmailResponse{
		Email: req.GetEmail(),
		Body:  data.Body,
	}, nil
}

func (s *Server) VerifyEmail(ctx context.Context, req *pbe.VerifyEmailRequest) (*pbe.VerifyEmailResponse, error) {
	verificationCode, err := s.GetData(req.GetEmail())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Verification code not found")
	} else {
		err := s.DeleteData(req.GetEmail())
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Failed to Delete Verification Code")
		}
	}

	if req.GetVerificationCode() == verificationCode {
		createSessionidRep, err := s.CreationSessionId(ctx, &pbs.CreateSessionIdRequest{Email: req.GetEmail(), Status: "SignUp"})
		if err != nil {
			return nil, err
		}

		return &pbe.VerifyEmailResponse{
			Token:   createSessionidRep.Token,
			Message: "Successfully created session",
		}, nil
	}

	return nil, status.Errorf(codes.Unauthenticated, "Verification code is incorrect")
}
