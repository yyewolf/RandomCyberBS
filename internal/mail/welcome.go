package mail

import (
	"fmt"
	"rcbs/internal/env"
	"rcbs/models"

	"github.com/nilslice/email"
	"github.com/sirupsen/logrus"
)

func SendWelcomeMail(u *models.User) {
	c := env.Get()

	logrus.WithFields(logrus.Fields{
		"username": u.Username,
		"email":    u.EmailAddress,
	}).Info("Sending welcome email")

	msg := email.Message{
		To:      u.EmailAddress,
		From:    fmt.Sprintf("no-reply@%s", c.Server.MailDomain),
		Subject: "RCBS - Welcome!",
		Body: fmt.Sprintf("Welcome to RCBS, %s!\n\n", u.Username) +
			"Thank you for signing up to RCBS. We hope you enjoy your time here.\n" +
			"To verify your email address, please click the link below:\n" +
			fmt.Sprintf("%s/api/v1/users/%s/verify/%s\n\n", c.Server.BaseURI, u.ID, u.VerificationToken) +
			"If you have any questions or need help, please don't hesitate to contact us.\n\n" +
			"Best regards,\n" +
			"The RCBS Team",
	}

	err := msg.Send()
	if err != nil {
		logrus.WithError(err).Error("failed to send welcome email")
	}

	logrus.WithFields(logrus.Fields{
		"username": u.Username,
		"email":    u.EmailAddress,
	}).Info("Welcome email sent")
}
