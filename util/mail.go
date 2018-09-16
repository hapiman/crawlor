package util

import (
	"fmt"
	"net/smtp"
	"strings"

	"github.com/hapiman/crawlor/config"
)

type unencryptedAuth struct {
	smtp.Auth
}

func (a unencryptedAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	s := *server
	s.TLS = true
	return a.Auth.Start(&s)
}

func sendMail(user, password, host, to, subject, body, mailtype string) error {
	// fmt.Printf("user=%s, password=%s, host=%s, to=%s, subject=%s, body=%s, mailtype=%s \n", user, password, host, to, subject, body, mailtype)
	auth := unencryptedAuth{
		smtp.PlainAuth(
			"",
			user,
			password,
			host,
		),
	}
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(config.STMP_HOST+":"+config.STMP_PORT, auth, user, send_to, msg)
	return err
}

func SendMsgToEmail(subject, msg string) error {
	err := sendMail(config.STMP_USER, config.STMP_PASSWORD, config.STMP_HOST, config.STMP_RECEIVER, subject, msg, "html")
	if err != nil {
		fmt.Println(err)
	}
	return err
}
