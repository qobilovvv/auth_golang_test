package tests

import (
	"testing"

	"github.com/qobilovvv/test_tasks/auth/internal/repositories"
	"github.com/qobilovvv/test_tasks/auth/internal/services"
)

func newOtpService() services.OTPService {
	repo := repositories.NewOTPRepository(testDB)
	return services.NewOTPService(repo)
}

func TestSendOtp(t *testing.T) {
	// service := newOtpService()

	// _, err := service.SendOTP("qobilovodiljonoff@gmail.com")
	// if err != nil {
	// 	t.Fatal("something went wrong: ", err)
	// }
}
