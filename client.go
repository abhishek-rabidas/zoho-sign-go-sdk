package zohosign

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/charmbracelet/log"
	"io"
	"net/http"
	"strings"
)

const (
	baseURL = "https://sign.zoho.in/api/v1"
)

type ZohoSignClient struct {
	authToken    string
	refreshToken string
	clientId     string
	secretKey    string
}

func NewZohoSignClient(refreshToken, clientId, secretKey string, refreshTokenValidMinutes int) *ZohoSignClient {
	client := ZohoSignClient{
		refreshToken: refreshToken,
		clientId:     clientId,
		secretKey:    secretKey,
	}

	authToken, err := generateAccessToken(client)

	if err != nil {
		log.Error("Zoho Sign Client ERROR", "Failed to generate Auth token", err.Error())
		return nil
	}

	client.authToken = authToken

	go client.refreshAuthToken(refreshTokenValidMinutes)
	return &client
}

func (z *ZohoSignClient) refreshAuthToken(timer int) {

}

func (z *ZohoSignClient) CreateTemplateSignRequest(recipientName, email, phoneNumber, countryCode, templateID, actionId, role, notes string, isEmbedded bool) (string, string, error) {
	templateAction := TemplateAction{
		RecipientName:   recipientName,
		RecipientEmail:  email,
		ActionId:        actionId,
		SigningOrder:    1,
		Role:            role,
		VerifyRecipient: false,
		PrivateNotes:    "",
		IsEmbedded:      isEmbedded,
	}

	if strings.TrimSpace(phoneNumber) != "" && strings.TrimSpace(countryCode) != "" {
		templateAction.RecipientPhonenumber = phoneNumber
		templateAction.RecipientCountrycode = countryCode
	}

	template := TemplateSignature{
		Notes:   notes,
		Actions: []TemplateAction{templateAction},
	}

	signRequest := SendTemplateSignatureRequest{
		Templates: template,
	}

	output, err := z.sendTemplateSignRequest(templateID, signRequest)
	if err != nil {
		log.Error("Error sending sign request", "err", err)
		return "", "", errors.New("Error sending sign request")
	}
	return output.Requests.RequestId, output.Requests.Actions[0].ActionId, nil
}

func (z *ZohoSignClient) sendTemplateSignRequest(templateId string, request SendTemplateSignatureRequest) (TemplateSignatureResponse, error) {
	log.Info("Sending sign request", "templateId", templateId)
	data, err := z.post(fmt.Sprintf("/templates/%s/createdocument?is_quicksend=%s", templateId, "true"), request)
	if err != nil {
		return TemplateSignatureResponse{}, err
	}
	var response TemplateSignatureResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return TemplateSignatureResponse{}, err
	}
	return response, nil
}

func (z *ZohoSignClient) post(path string, request any) ([]byte, error) {
	body, err := json.Marshal(request)

	if err != nil {
		return nil, err
	}

	log.Debug("Request Body", string(body))

	r, err := http.NewRequest("POST", fmt.Sprintf("%s%s", baseURL, path), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", fmt.Sprintf("Zoho-oauthtoken %s", z.authToken))

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}

	data, _ := io.ReadAll(res.Body)

	log.Debug("Signature Request API Response", string(data))

	return data, nil
}

func (z *ZohoSignClient) call(path string, method string) ([]byte, error) {
	r, err := http.NewRequest(method, fmt.Sprintf("%s%s", baseURL, path), nil)
	if err != nil {
		return nil, err
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", fmt.Sprintf("Zoho-oauthtoken %s", z.authToken))

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}

	data, _ := io.ReadAll(res.Body)

	log.Debug("Signature Request API Response", string(data))

	return data, nil
}

func (z *ZohoSignClient) CheckDocumentStatus(requestId string) (TemplateSignatureResponse, error) {
	var response TemplateSignatureResponse

	data, err := z.call(fmt.Sprintf("/requests/%s", requestId), "GET")

	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, errors.New("Failed to fetch document status")
	}

	return response, nil
}

func (z *ZohoSignClient) GetEmbeddedSignatureURL(requestId, documentActionId string) (EmbeddedSignedDocumentResponse, error) {
	var response EmbeddedSignedDocumentResponse

	data, err := z.call(fmt.Sprintf("/requests/%s/actions/%s/embedtoken", requestId, documentActionId), "POST")

	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, errors.New("Failed to fetch document signed url")
	}

	return response, nil
}

func (z *ZohoSignClient) CancelSignatureRequest(requestId string) error {
	_, err := z.call(fmt.Sprintf("/requests/%s/recall", requestId), "POST")
	return err
}

func (z *ZohoSignClient) UpdateAuthToken(token string) {
	z.authToken = token
}

func generateAccessToken(z ZohoSignClient) (string, error) {
	r, err := http.NewRequest("POST", fmt.Sprintf("https://accounts.zoho.in/oauth/v2/token?refresh_token=%s&client_id=%s&client_secret=%s&grant_type=refresh_token",
		z.refreshToken, z.clientId, z.secretKey), nil)
	if err != nil {
		return "", err
	}

	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	data, _ := io.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK {
		log.Error("Access token generation failed", "res", string(data))
		return "", errors.New(res.Status)
	}

	var authResponse AuthRefreshResponse

	err = json.Unmarshal(data, &authResponse)

	if err != nil {
		return "", err
	}

	return authResponse.AccessToken, nil
}
