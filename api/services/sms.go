// sms.go

package services

import (
	"log"
	"os"

	twilio "github.com/kevinburke/twilio-go"
)

// Sms send message via Twilio
func Sms(params map[string]string) {

	client := twilio.NewClient(os.Getenv("TWILIO_SID"), os.Getenv("TWILIO_TOKEN"), nil)

	log.Println("SID", os.Getenv("TWILIO_SID"))
	log.Println("TOKEN", os.Getenv("TWILIO_TOKEN"))

	success, err := client.Messages.SendMessage(params["from"], params["to"], params["message"], nil)

	log.Println("success", success)
	log.Println("err", err)
}
