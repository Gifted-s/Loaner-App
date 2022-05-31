package helpers

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"
)
func SendMail(name string,email string, subject string,htmlContent string) {

	from := mail.NewEmail("Helpful","sunkanmiadewumi1@gmail.com")
	to := mail.NewEmail(name,email)
	message := mail.NewSingleEmail(from, subject, to, "", htmlContent)
	client := sendgrid.NewSendClient("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")
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}