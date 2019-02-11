// mail.go

package services

import (
	"os"

	h "github.com/jmilagroso/api/helpers"
	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// Mail sends email
func Mail(params map[string]string) {
	from := mail.NewEmail(params["from_name"], params["from_email"])
	subject := params["subject"]
	to := mail.NewEmail(params["to_name"], params["to_email"])
	plainTextContent := params["content_plaintext"]
	htmlContent := params["content_html"]
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

	_, sendError := client.Send(message)

	h.Error(sendError)
}
