package main

import (
    "fmt"
    "os"
    gomail "gopkg.in/mail.v2"
)

func main() {
	if  len(os.Args) < 1 {
		fmt.Println("You have to pass the mail message and file path to attach")
		os.Exit(2)
	}
	subject := os.Args[1]
	message := os.Args[2]
	fileToAttachPath := os.Args[3]
  sendEmail(subject, message, fileToAttachPath)
}

func sendEmail(subject string, message string, filePath string){
	// Sender data.
  userId := os.Getenv("SMTP_USERNAME")
  from := os.Getenv("MAIL_FROM")
  toEmail := os.Getenv("TO_EMAIL")
  password :=  os.Getenv("SMTP_PASSWORD")
  smtpHost := os.Getenv("SMTP_HOST")
  
   m := gomail.NewMessage()

  // Set E-Mail sender
  m.SetHeader("From", from)

  // Set E-Mail receivers
  m.SetHeader("To", toEmail)

  // Set E-Mail subject
  m.SetHeader("Subject", subject)

  // Set E-Mail body. You can set plain text or html with text/html
  m.SetBody("text/plain", message)

  m.Attach(filePath)

  // Settings for SMTP server
  d := gomail.NewDialer(smtpHost, 465, userId, password)

  // This is only needed when SSL/TLS certificate is not valid on server.
  // In production this should be set to false.
  // d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

  // Now send E-Mail
  if err := d.DialAndSend(m); err != nil {
    fmt.Println(err)
    panic(err)
  }
}
