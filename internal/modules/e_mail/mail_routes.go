package email

import (

	"github.com/gofiber/fiber/v2"
)

func EmailRoutes(r fiber.Router) {

	r.Post("/sendotp", SentOTPController)
	r.Post("/verifyotp", VerifyOTPController)
	r.Post("/sendemail", EmailContactController)
}