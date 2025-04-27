package Controllers

import (
	database "Go-Auth/Database"
	"Go-Auth/model"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

const secretKey = "secret"

func Register(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var existingUser model.User
	database.DB.Where("LOWER(email) = ?", data["email"]).First(&existingUser)
	if existingUser.ID != 0 { // Assuming ID is the primary key
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email is already registered",
		})
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)
	user := model.User{
		Name:       data["name"],
		Email:      data["email"],
		Password:   password,
		IsVerified: true, // 🔥 Default false
	}

	database.DB.Create(&user)

	//sendVerificationEmail(user.Email)
	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var user model.User
	database.DB.Where("email = ?", data["email"]).First(&user)
	if user.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User not found",
		})
	}
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	// Example user ID (replace with your actual user.ID)
	userID := int64(user.ID)

	// Set expiration time
	expireTime := time.Now().Add(24 * time.Hour)

	// Create the claims
	claims := &jwt.StandardClaims{
		Issuer:    strconv.FormatInt(userID, 10),           // Converts userID to string
		ExpiresAt: jwt.NewTime(float64(expireTime.Unix())), // Adjusted for jwt-go/v4
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with a secret
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Cloud not sign token",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    signedToken,
		Expires:  expireTime,
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "Login successful",
	})
}

func UserDetails(c *fiber.Ctx) error {
	// Get the JWT token from cookies
	cookie := c.Cookies("jwt")
	if cookie == "" {
		return c.Status(401).JSON(fiber.Map{
			"message": "Missing or invalid cookie",
		})
	}

	// Parse the JWT token
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil || token == nil || !token.Valid {
		return c.Status(401).JSON(fiber.Map{
			"message": "UnAuthenticated",
		})
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to parse claims",
		})
	}

	var user model.User
	database.DB.Where("id = ?", claims.Issuer).First(&user)

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "Logout successful",
	})
}
