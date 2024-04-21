package main

import (
	"flag"
	"log"
	"net/smtp"
	"os"
	"path/filepath"
)

func main() {
	// Command-line flags
	var (
		name       string
		subject    string
		body       string
		attachment string
	)

	// Parse command-line flags
	flag.StringVar(&name, "name", "", "Sender's name")
	flag.StringVar(&subject, "subject", "", "Email subject")
	flag.StringVar(&body, "body", "", "Email body")
	flag.StringVar(&attachment, "attachment", ".", "Attachment file or directory path")
	flag.Parse()

	// Validate required flags
	if name == "" || subject == "" || body == "" {
		log.Fatal("Missing required flags. Usage: go run main.go -name=<name> -subject=<subject> -body=<body> [-attachment=<attachment>]")
	}

	// Sender's email information
	from := "sender@example.com"
	smtpServer := "smtp.example.com"
	smtpPort := "587"
	smtpUsername := "sender@example.com"
	smtpPassword := "password"

	// Recipient's email information
	to := "recipient@example.com"

	// Set up authentication for SMTP server
	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpServer)

	// Attachments
	var attachments []string
	if attachment != "." {
		files, err := os.ReadDir(attachment)
		if err != nil {
			log.Fatal("Error reading directory:", err)
		}

		// Collect file paths for attachments
		for _, file := range files {
			if file.IsDir() {
				continue // Skip directories
			}
			attachments = append(attachments, filepath.Join(attachment, file.Name()))
		}
	}

	// Send the email
	err := SendEmail(from, to, smtpServer+":"+smtpPort, auth, name, subject, body, attachments)
	if err != nil {
		log.Fatal("Error sending email:", err)
	}

	log.Println("Email sent successfully!")
}

// SendEmail sends an email with attachments
func SendEmail(from, to, smtpServer string, auth smtp.Auth, name, subject, body string, attachments []string) error {
	// Compose the email message
	message := []byte(
		"From: " + name + " <" + from + ">\r\n" +
			"To: " + to + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: multipart/mixed; boundary=boundary123456\r\n\r\n" +
			"--boundary123456\r\n" +
			"Content-Type: text/plain; charset=UTF-8\r\n\r\n" +
			body + "\r\n\r\n",
	)

	// Attach files
	for _, file := range attachments {
		content, err := os.ReadFile(file)
		if err != nil {
			return err
		}
		filename := filepath.Base(file)
		message = append(message,
			[]byte(
				"--boundary123456\r\n"+
					"Content-Type: application/octet-stream\r\n"+
					"Content-Disposition: attachment; filename=\""+filename+"\"\r\n"+
					"Content-Transfer-Encoding: base64\r\n\r\n"+
					string(content)+"\r\n\r\n",
			)...,
		)
	}

	// Send the email using SMTP
	return smtp.SendMail(smtpServer, auth, from, []string{to}, message)
}
