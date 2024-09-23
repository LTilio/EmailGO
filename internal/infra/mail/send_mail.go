package mail

import (
	"EmailGO/internal/domain/campaign"
	"fmt"
	"os"
	"time"

	"gopkg.in/gomail.v2"
)

func SendMail(campaign *campaign.Campaign) error {
	fmt.Println("Sending mail...")

	// start := time.Now() //metodo para marcar o tempo
	d := gomail.NewDialer(os.Getenv("EMAIL_SMTP"), 587, os.Getenv("EMAIL_USER"), os.Getenv("EMAIL_PASSWORD"))
	// duration := time.Since(start) //metodo para ver o tempo de duração do metodo
	// fmt.Println("dialer crated in ", duration)

	var emails []string
	for _, contacs := range campaign.Contacts {
		emails = append(emails, contacs.Email)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL_USER"))
	m.SetHeader("To", emails...)
	m.SetHeader("Subject", campaign.Name)
	m.SetBody("text/html", campaign.Content)

	start := time.Now()
	err := d.DialAndSend(m)
	duration := time.Since(start)
	fmt.Println("DialAndSend in ", duration)

	return err
}
