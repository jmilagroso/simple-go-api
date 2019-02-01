// mail.go
// Mail endpoints
// Jay Milagroso <jmilagroso@quadx.xyz> / Jan 24 2019

package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"quadx.xyz/jmilagroso/goberks/blueprints"
	h "quadx.xyz/jmilagroso/goberks/helpers"
	s "quadx.xyz/jmilagroso/goberks/services"
)

// MailHook Hook local type
type MailHook blueprints.Hook

// Notify endpoint
func Notify(w http.ResponseWriter, r *http.Request) {

	reqBody, readAllError := ioutil.ReadAll(r.Body)
	h.Error(readAllError)

	hook := MailHook{}
	h.Error(json.Unmarshal([]byte(reqBody), &hook))

	emailParams := make(map[string]string)
	emailParams["subject"] = "QUADX Delivery"
	emailParams["from_name"] = "Quadx"
	emailParams["from_email"] = "no-reply@quadx.xyz"
	emailParams["to_name"] = "Jay"
	emailParams["to_email"] = "j.milagroso@gmail.com" // @TODO Get from current user
	emailParams["content_plaintext"] = "Hello Jay, We've detected that your package with Tracking Number XXXX-XXXX-XXXX is nearby! Please prepare your payment."
	emailParams["content_html"] = "Hello Jay, <br><br>We've detected that your package with <strong>Tracking Number XXXX-XXXX-XXXX</strong> is nearby! <br><br>Please prepare your payment."
	s.Mail(emailParams)

	smsParams := make(map[string]string)
	smsParams["from"] = "+639156589346"
	smsParams["to"] = "+639156589346" // @TODO Get from current user
	smsParams["message"] = "Your package develivery is on the way! Tracking Number XXXX-XXXX-XXXX. QuadX/XPOST"
	s.Sms(smsParams)
}
