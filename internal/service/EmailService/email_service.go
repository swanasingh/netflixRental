package EmailService

import (
	"fmt"
	"net/smtp"
	"netflixRental/configs"
)

type EmailService interface {
	GetEmailConfig()
	GetAuth() smtp.Auth
	SendInvoice(toEmail string, body string)
}

type emailService struct {
	configs.EmailConfig
}

func (e *emailService) GetEmailConfig() {
	fmt.Println("get email config")
	var config = configs.Config{}
	configs.GetConfigs(&config)
	e.Host = config.Email.Host
	e.Port = config.Email.Port
	e.FromEmail = config.Email.FromEmail
	e.Password = config.Email.Password
}

func (e emailService) GetAuth() smtp.Auth {
	auth := smtp.PlainAuth("", e.FromEmail, e.Password, e.Host)
	return auth

}

func (e emailService) SendInvoice(toEmail string, body string) {
	auth := e.GetAuth()
	err := smtp.SendMail(e.Host+":"+e.Port, auth, e.FromEmail, []string{toEmail}, []byte(body))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Email sent successfully!")
}

func NewEmailService() EmailService {
	var emailSer emailService
	emailSer.GetEmailConfig()
	fmt.Println("got email config")
	return &emailSer
}
