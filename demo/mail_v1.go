package main

import "net/smtp"

const DISCLAIMEER = `免责声明，此电子邮件和任何附件可能包含特权和机密信息，
仅供指定的收件人使用。如果你错误收到此电子邮件，请发立即删除电子邮件，
任何保密性的等等，那就不管了`

func SendMailWithDisclaimer(subject, from string, to []string, content string,
	mailServer string, a smtp.Auth) error {
	e := smtp.NewEmail{}
	e.From = from
	e.To = to
	e.Subject = subject
	e.Text = []byte(attachDisclaimer(content))
	return e.Send(mailServer, a)

}
