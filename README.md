# Zoho Sign Go SDK

A simple and lightweight Go SDK for interacting with the [Zoho Sign API](https://www.zoho.com/sign/api/). This SDK makes it easy to send signature requests using templates, manage embedded signing, check document status, and cancel signature requests via Zoho Sign.

## Features

- ✅ Generate and auto-refresh OAuth access tokens
- 📨 Send template-based signature requests
- 🔍 Check document request status
- 🔗 Retrieve embedded signing URL
- ❌ Cancel signature requests

## Installation

```bash
go get github.com/abhishek-rabidas/zoho-sign-go-sdk
```
## Usage

### 1. Import and Initialize

```go
import zohosign "github.com/abhishek-rabidas/zoho-sign-go-sdk"

zohoSignClient := zohosign.NewZohoSignClient(
    "your_refresh_token", 
    "your_client_id", 
    "your_secret_key", 
    5, // Refresh token interval in minutes
)
```
### 2. Send Signature Request

```go
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
```
### 3. Check Document Status

```go
docResponse, err := zohoSignClient.CheckDocumentStatus(res.Requests.RequestId)
if err != nil {
    log.Error(err.Error())
    return
}
log.Infof("Document Status: %+v", docResponse)
```

### 4. Cancel Signature Request

```go
err = zohoSignClient.CancelSignatureRequest(res.Requests.RequestId)
if err != nil {
    log.Error(err.Error())
    return
}
log.Info("Document cancelled successfully")
```

### 5. Get Embedded Signing URL

```go
embeddedSignResponse, err := zohoSignClient.GetEmbeddedSignatureURL(res.Requests.RequestId, res.Requests.Actions[0].ActionId)
if err != nil {
    log.Error(err.Error())
    return
}
log.Infof("Embedded Signature URL: %s", embeddedSignResponse.SignUrl)


```