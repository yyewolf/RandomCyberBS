package mail

import (
	"crypto/tls"
	"fmt"
	"rcbs/internal/env"
	"rcbs/models"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

func SendWelcomeMail(u *models.User) {
	c := env.Get()

	var m gomail.Message

	m.SetAddressHeader("From", fmt.Sprintf("no-reply@%s", c.Server.Mail.From), c.Server.Mail.Name)
	m.SetHeader("To", u.EmailAddress)
	m.SetHeader("Date", m.FormatDate(time.Now()))
	m.SetHeader("Subject", c.Server.Mail.Name+" - Welcome!")
	m.SetBody("text/plain",
		fmt.Sprintf("Welcome to RCBS, %s!\n\n", u.Username)+
			"Thank you for signing up to RCBS. We hope you enjoy your time here.\n"+
			"To verify your email address, please click the link below:\n"+
			fmt.Sprintf("%s/api/v1/users/%s/verify/%s\n\n", c.Server.BaseURI, u.ID, u.VerificationToken)+
			"If you have any questions or need help, please don't hesitate to contact us.\n\n"+
			"Best regards,\n"+
			"The RCBS Team")

	logrus.WithFields(logrus.Fields{
		"username": u.Username,
		"email":    u.EmailAddress,
	}).Debug("Sending welcome email")

	d := gomail.NewDialer(c.Server.Mail.RelayHost, c.Server.Mail.RelayPort, c.Server.Mail.RelayUsername, c.Server.Mail.RelayPassword)
	if c.Server.Mail.RelayTLS {
		d.TLSConfig = &tls.Config{InsecureSkipVerify: c.Server.Mail.RelayIgnoreTLSCert}
	}

	d.DialAndSend()

	logrus.WithFields(logrus.Fields{
		"username": u.Username,
		"email":    u.EmailAddress,
	}).Debug("Welcome email sent")
}
