package email

import (
	"navora/packages/utils"

	"github.com/gofiber/fiber/v2"
)

func EmailContactController(c *fiber.Ctx) error {
	var req ContactRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.Error(c, 400, "Invalid request", err.Error())
	}

	if req.Name == "" || req.Email == "" || req.Message == "" {
		return utils.Error(c, 400, "Please fill all required fields", nil)
	}

	message, err := EmailContactService(req)
	if err != nil {
		return utils.Error(c, 500, message, err.Error())
	}

	return utils.Success(c, 200, "Message sent successfully", nil)
}

func SentOTPController(c *fiber.Ctx) error {
	var OtpEmail struct {
		Email string `json:"email" validate:"required,email"`
	}

	if err := c.BodyParser(&OtpEmail); err != nil {
		return utils.Error(c, 400, "invalid request", nil)
	}

	if OtpEmail.Email == "" {
		return utils.Error(c, 400, "email required", nil)
	}

	otp, err := SentOTPService(OtpEmail.Email)
	if err != nil {
		return utils.Error(c, 500, "failed to send the OTP", err.Error())
	}

	return utils.Success(c, 200, "OTP sented", otp)
}

func VerifyOTPController(c *fiber.Ctx) error {
	var VerifyOtp struct {
		Email string `json:"email" validate:"required,email"`
		Otp   string `json:"otp" validate:"required,len=6,numeric"`
	}

	if err := c.BodyParser(&VerifyOtp); err != nil {
		return utils.Error(c, 400, "Invalied Input", err.Error())
	}

	if err := VerifyOTPService(VerifyOtp.Email, VerifyOtp.Otp); err != nil {
		return utils.Error(c, 500, "Email Verification Failed", err.Error())
	}

	return utils.Success(c, 200, "Email Verified Successfully", nil)
}