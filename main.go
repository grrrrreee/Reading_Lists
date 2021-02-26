package main

import (
	"log"
	"net/smtp"
)

func main() {
	to_email := "parkms940@naver.com"
	from_email := "greedycatty@gmail.com"
	password := "r1s1e1!1"
	subject_body := "Subject: Write your Subject \n\n" + "This is your Email body"
	status := smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", from_email, password, "smtp.gmail.com"), from_email, []string{to_email}, []byte(subject_body))
	if status != nil {
		log.Printf("Error from SMTP Server: %s", status)
	}
	log.Print("Email Sent Successfully")
}
