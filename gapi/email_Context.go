package gapi

import (
	"context"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pbu "Email/pb/users"
	"Email/util"
)

type EmailContext struct {
	Subject string
	Body    string
	Email   []string
}

func (s *Server) EmailContextTpye(number, email string) (EmailContext, error) {
	switch number {
	case "signup":
		verificationCode, err := util.GenerateVerificationCode()
		if err != nil {
			return EmailContext{}, status.Error(codes.Internal, "Failed to Create Verification Code")
		}

		body := `
			<html>
			<body>
				<div>
					<p>親愛的用戶，</p>
					<p>感謝您註冊我們的服務！為了完成註冊流程，請使用以下驗證碼來驗證您的帳戶：</p>
					<p style="font-size: 18px; font-weight: bold">驗證碼：` + verificationCode + `</p>
					<p>此驗證碼有效期為 5 分鐘。如果您沒有進行註冊操作，請忽略此郵件。</p>
					<p>若您有任何問題，請隨時聯繫我們的客服團隊。</p>
					<p>謝謝！</p>
					<p>祝您使用愉快，<br>Cyrus Man</p>
				</div>
			</body>
			</html>
			`

		err = s.SaveVerificationCode(email, verificationCode)
		if err != nil {
			return EmailContext{}, status.Error(codes.Internal, "Failed to save the verification code in Redis")
		}

		return EmailContext{
			Subject: "XXX - 註冊驗證碼",
			Body:    body,
			Email:   []string{email},
		}, nil

	case "after_login_reset_password":
		// TODO reset password url setting
		user, err := s.UserInformations(context.Background(), &pbu.UserInformationRequest{Email: email})
		if err != nil {
			return EmailContext{}, status.Error(codes.Internal, "Failed to get user information")
		}

		body := `
		<html>
		<body>
			<div>
				<p>親愛的 ` + user.FristName + user.LastName + `，</p> 
				<p>您好！</p>
				<p>我們收到您重設帳戶密碼的請求。請點擊下方連結以設定新的密碼：</p>
				<a href="http://localhost:8080" style="font-size: 18px; font-weight: bold; padding: 10px 20px; border-radius: 5px;">
					重設密碼連結
				</a>
				<p>若您未發出此請求，請忽略此郵件或聯繫我們的客服團隊以確保您的帳戶安全。</p>
				<p>若您有任何問題，請隨時聯繫我們的客服團隊。</p>
				<p>謝謝！</p>
				<p>祝您使用愉快</p>
				<p>Cyrus Man</p>
			</div>
		</body>
		</html>
		`
		return EmailContext{
			Subject: "XXX - 重設密碼",
			Body:    body,
			Email:   []string{email},
		}, nil

	case "reset Done":

	}

	return EmailContext{}, nil
}

func (s *Server) SaveVerificationCode(email, verificationCode string) error {
	err := s.SetData(email, verificationCode, 5*time.Minute)
	if err != nil {
		return status.Errorf(codes.Internal, "Fail to Save to Redis")
	}

	return nil
}
