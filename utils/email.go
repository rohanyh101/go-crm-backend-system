package utils

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/rohanhonnakatti/golang-jwt-auth/models"
)

func SendEmailInteraction(interaction models.Interaction, email string) {
	// Example email content
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	emailFrom := os.Getenv("SMTP_MAIL")
	emailPassword := os.Getenv("SMTP_PASSWORD")
	// to

	subject := "Ticket Created: " + *interaction.Title
	body := fmt.Sprintf("Dear User,\n\nYour ticket with ID %s has been created.\n\nDetails:\nDescription: %s\nStatus: %s\n\nThank you,\nSupport Team", *interaction.Title, *interaction.Description)
	to := []string{email}

	auth := smtp.PlainAuth("", emailFrom, emailPassword, smtpHost)

	msg := []byte("To: " + to[0] + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	// Send the email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, emailFrom, to, msg)
	if err != nil {
		fmt.Println("Error sending email:", err)
		return
	}

	fmt.Println("Email sent successfully to:", to)
}

func SendEmailToUser(ticket models.Ticket, email string) {
	// Example email content
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	emailFrom := os.Getenv("SMTP_MAIL")
	emailPassword := os.Getenv("SMTP_PASSWORD")
	// to

	subject := "Ticket Created: " + ticket.TicketId
	body := fmt.Sprintf("Dear User,\n\nYour ticket with ID %s has been created.\n\nDetails:\nDescription: %s\nStatus: %s\n\nThank you,\nSupport Team", ticket.TicketId, *ticket.Description, *ticket.Status)
	to := []string{email}

	// Set up authentication information.
	auth := smtp.PlainAuth("", emailFrom, emailPassword, smtpHost)

	// Compose the email message
	msg := []byte("To: " + to[0] + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	// Send the email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, emailFrom, to, msg)
	if err != nil {
		fmt.Println("Error sending email:", err)
		return
	}

	fmt.Println("Email sent successfully to:", to)
}
