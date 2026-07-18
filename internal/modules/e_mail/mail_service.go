package email

import (
	"errors"
	"fmt"
	"log"
	"navora/packages/utils"
	"os"
)

func EmailContactService(req ContactRequest) (string, error) {

	companyhtml := fmt.Sprintf(`
		<h2>New Contact Request</h2>

		<p><strong>Name:</strong> %s</p>
		<p><strong>Email:</strong> %s</p>
		<p><strong>Subject:</strong> %s</p>

		<p><strong>Message:</strong></p>

		<p>%s</p>
	`,
		req.Name,
		req.Email,
		req.Subject,
		req.Message,
	)

	err := utils.SentToEmail(os.Getenv("COMPANY_EMAIL"), "New Contact Request", companyhtml)
	if err != nil {
		return "Unable to send email", err
	}

	userHTML := fmt.Sprintf(`
		<div style="font-family:Arial,sans-serif;max-width:600px;margin:auto;padding:20px">

			<h2>Thank You for Contacting Navora Technologies</h2>

			<p>Hi <strong>%s</strong>,</p>

			<p>Thank you for contacting <strong>Navora Technologies</strong>.</p>

			<p>We have successfully received your request.</p>

			<p>Our team will review your message and get back to you within <strong>24 hours</strong>.</p>

			<hr>

			<p><strong>Your Request</strong></p>

			<p><b>Subject:</b> %s</p>

			<p>%s</p>

			<br>

			<p>Best Regards,</p>
			<p><strong>Navora Technologies</strong></p>

		</div>
	`,
		req.Name,
		req.Subject,
		req.Message,
	)

	if err := utils.SentToEmail(req.Email, "Thank You for Contacting Navora Technologies", userHTML,); err != nil {
		log.Printf("failed to send thank-you email: %v", err)
	}

	return "Message sent successfully", nil
}

func SentOTPService(email string) (otp string, err error) {
	isOk, err := utils.RateLimitOTP(email)

	if err != nil {
		return "", err
	}

	if !isOk {
		return "", errors.New("Request limit exceeded, wait for 5 min")
	}

	RandOTP, errOTP := utils.GenerateOTP()
	if errOTP != nil {
		return "",  errOTP
	}
	fmt.Println("OTP :", RandOTP)

	htmlContent := "<h2>Your OTP</h2><b>" + RandOTP + "</b><p>Valid for 5 min</p>"

	if err := utils.SentToEmail(email, "Your Navora OTP", htmlContent); err != nil {
		return "", err
	}

	if err := utils.SaveOTP(email, RandOTP); err != nil {
		return "", errors.Join(errors.New("failed to save the OTP"), err)
	}

	return RandOTP, nil
}

func VerifyOTPService(email, otp string) error {
	storedOTP, err := utils.GetOTP(email)

	if err != nil {
		return errors.Join(errors.New("OTP mismatched"), err)
	}

	if storedOTP != otp {
		return errors.New("invalid otp")
	}

	utils.DeleteOTP(email)
	return nil
}
