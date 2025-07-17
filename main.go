package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gopkg.in/gomail.v2"
)

// We create a struct to hold the data from the incoming JSON request.
// The `json:"..."` part tells Go how to map the JSON keys to our struct fields.
type ContactForm struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func main() {
	// Create a new Fiber app. Think of this as the main engine for our API.
	app := fiber.New()
	app.Use(cors.New())
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))

	// Define the POST route. This is the endpoint our front end will call.
	app.Post("/api/sendmail/", sendMailHandler)

	// Start the server on port 3003.
	// You can access this at http://localhost:3003
	log.Fatal(app.Listen(":3003"))
}

// This is the function that handles the request when someone hits our endpoint.
func sendMailHandler(c *fiber.Ctx) error {
	// 1. PARSE THE INCOMING REQUEST
	// ---------------------------------
	// Create an empty ContactForm to store the parsed data.
	form := new(ContactForm)

	// Parse the request body (the JSON data) into our 'form' struct.
	// If there's an error (like malformed JSON), we send a 400 Bad Request response.
	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// 2. CONFIGURE AND SEND THE EMAIL
	// ---------------------------------
	// Get email credentials from environment variables.
	// IMPORTANT: Never hardcode go run main.goords in your code!
	emailUser := os.Getenv("EMAIL_USER")
	emailPass := os.Getenv("EMAIL_PASS")

	// Check if the environment variables are set.
	if emailUser == "" || emailPass == "" {
		log.Println("Email credentials are not set in environment variables.")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Email server not configured.",
		})
	}

	// Create a new email message.
	m := gomail.NewMessage()

	// Set the sender and recipient.
	m.SetHeader("From", emailUser)
	m.SetHeader("To", emailUser) // Sending to yourself

	// Set the email subject.
	m.SetHeader("Subject", "You've got a mail, bro.")

	// Construct the email body from the form data.
	// This makes the email content dynamic based on the user's input.
	emailBody := fmt.Sprintf(
		"You received a new message:\n\nName: %s\nEmail: %s\nTitle: %s\n\nContent:\n%s",
		form.Name, form.Email, form.Title, form.Content,
	)
	m.SetBody("text/plain", emailBody)

	// Configure the SMTP server connection.
	// We're using Yahoo's SMTP server details here.
	// Port 587 is standard for SMTP with TLS encryption.
	d := gomail.NewDialer("smtp.mail.yahoo.com", 587, emailUser, emailPass)

	// Send the email.
	// If it fails, we log the error and send a 500 Internal Server Error response.
	if err := d.DialAndSend(m); err != nil {
		log.Println("Could not send email: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to send email.",
		})
	}

	// 3. SEND A SUCCESS RESPONSE
	// ---------------------------------
	// If everything went well, send a 200 OK response with a success message.
	log.Println("Email sent successfully!")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Email sent successfully!",
	})
}
