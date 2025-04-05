package main

import (
	zohosign "github.com/abhishek-rabidas/zoho-sign-go-sdk"
	"github.com/charmbracelet/log"
)

func main() {
	zohoSignClient := zohosign.NewZohoSignClient("your_refresh_token", "your_client_id", "your_secret_key", 5)

	// Sending a template signature request
	res, err := zohoSignClient.CreateTemplateSignRequest(
		"recipient_name",
		"recipient_email",
		"recipient_phone_number",
		"recipient_country_code",
		"template_id",
		"action_id",
		"role",
		"notes",
		true,
	)
	if err != nil {
		log.Error(err.Error())
		return
	}
	log.Infof("Template Signature Request Response: %+v", res)

	// Checking the status of a document
	docResponse, err := zohoSignClient.CheckDocumentStatus(res.Requests.RequestId)
	if err != nil {
		log.Error(err.Error())
		return
	}
	log.Infof("Document Status: %+v", docResponse)

	// Get Embedded Signature URL
	embeddedSignResponse, err := zohoSignClient.GetEmbeddedSignatureURL(res.Requests.RequestId, res.Requests.Actions[0].ActionId)
	if err != nil {
		log.Error(err.Error())
		return
	}
	log.Infof("Embedded Signature URL: %s", embeddedSignResponse.SignUrl)

	// Cancel a document
	err = zohoSignClient.CancelSignatureRequest(res.Requests.RequestId)
	if err != nil {
		log.Error(err.Error())
		return
	}
	log.Info("Document cancelled successfully")

}
