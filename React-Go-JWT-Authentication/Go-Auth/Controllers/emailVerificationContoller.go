package Controllers

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	database "Go-Auth/Database"
	"Go-Auth/model"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/gomail.v2"
)

var otpStore = make(map[string]string) // email -> otp

func sendVerificationEmail(email string) {
	rand.Seed(time.Now().UnixNano())
	otp := fmt.Sprintf("%06d", rand.Intn(1000000)) // generate random 6 digit otp
	otpStore[email] = otp

	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EmailVerificationID"))
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Your Verification Code")
	m.SetBody("text/plain", "Your verification code is: "+otp)
	var EmailVerificationID string = os.Getenv("EmailVerificationID")
	var EmailVerificationPassword string = os.Getenv("EmailVerificationPassword")
	d := gomail.NewDialer("smtp.gmail.com", 587, EmailVerificationID, EmailVerificationPassword)

	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Failed to send verification email:", err)
	}
}

func VerifyCode(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	email := data["email"]
	code := data["code"]

	expectedCode, exists := otpStore[email]
	if !exists || expectedCode != code {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid verification code",
		})
	}

	// Update user isVerified = true
	var user model.User
	database.DB.Where("LOWER(email) = ?", email).First(&user)
	if user.ID == 0 {
		return c.Status(400).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	user.IsVerified = true
	database.DB.Save(&user)

	// Remove OTP after verification
	delete(otpStore, email)

	return c.JSON(fiber.Map{
		"message": "Email verified successfully!",
	})
}
