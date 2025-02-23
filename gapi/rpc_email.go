package gapi

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"Email/pb"
	"Email/util"
)

type EmailContext struct {
	Subject string
	Body    string
	To      []string
}

func (s *Server) SendEmail(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	verificationCode, err := util.GenerateVerificationCode()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Fail to Create Verification Code")
	}

	body := `
	<html>
	<body>
		<div>
			<h2>註冊驗證碼</h2>
			<p>親愛的用戶，</p>
			<p>感謝您註冊我們的服務！為了完成註冊流程，請使用以下驗證碼來驗證您的帳戶：</p>
			<p style="font-size: 18px; font-weight: bold">驗證碼：` + verificationCode + `</p>
			<p>此驗證碼有效期為 10 分鐘。如果您沒有進行註冊操作，請忽略此郵件。</p>
			<p>若您有任何問題，請隨時聯繫我們的客服團隊。</p>
			<p>謝謝！</p>
			<p>祝您使用愉快，<br>Cyrus Man</p>
		</div>
	</body>
	</html>
	`

	data := EmailContext{
		Subject: "註冊驗證碼",
		Body:    body,
		To:      []string{req.GetTo()},
	}

	err = SendEmail(s.config, data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Fail to Send Email")
	}

	return &pb.SignUpResponse{Message: "Successfully sent the verification code email"}, nil
}
