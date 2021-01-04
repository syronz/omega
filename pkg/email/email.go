package email

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

// ConfigEmail is used for initiate the ConfigEmail design pattern
type ConfigEmail struct {
	From     string
	To       string
	Host     string
	Port     int
	Username string
	Password string
}

// EmailConfig initiate the functionality
func EmailConfig() *ConfigEmail {
	return &ConfigEmail{}
}

// Send email with body and attachment
func (c *ConfigEmail) SendEmail(subject, title, attachment string, body interface{}) {
	m := gomail.NewMessage()
	m.SetHeader("From", c.From)
	m.SetHeader("To", c.To)
	m.SetHeader("Subject", subject)
	if attachment != "" {
		m.Attach(attachment)
	}
	m.SetBody("text/html",
		fmt.Sprintf("<html><body><h4>%v</h4><br></body>"+
			"<pre>%v</pre>"+
			"</html>", title, body))

	d := gomail.NewPlainDialer(c.Host, c.Port, c.Username, c.Password)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
