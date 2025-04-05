package zohosign

// DocumentStatusResponse response for particular document signature status
type DocumentStatusResponse struct {
	Code     int     `json:"code"`
	Requests Request `json:"requests"`
	Message  string  `json:"message"`
	Status   string  `json:"status"`
}

type Request struct {
	RequestStatus       string              `json:"request_status"`
	Notes               string              `json:"notes"`
	Attachments         []interface{}       `json:"attachments"`
	ReminderPeriod      int                 `json:"reminder_period"`
	OwnerId             string              `json:"owner_id"`
	Description         string              `json:"description"`
	RequestName         string              `json:"request_name"`
	ModifiedTime        int64               `json:"modified_time"`
	ActionTime          int64               `json:"action_time"`
	IsDeleted           bool                `json:"is_deleted"`
	ExpirationDays      int                 `json:"expiration_days"`
	IsSequential        bool                `json:"is_sequential"`
	SignSubmittedTime   int64               `json:"sign_submitted_time"`
	TemplatesUsed       []string            `json:"templates_used"`
	OwnerFirstName      string              `json:"owner_first_name"`
	SignPercentage      float64             `json:"sign_percentage"`
	ExpireBy            int64               `json:"expire_by"`
	IsExpiring          bool                `json:"is_expiring"`
	OwnerEmail          string              `json:"owner_email"`
	CreatedTime         int64               `json:"created_time"`
	EmailReminders      bool                `json:"email_reminders"`
	DocumentIds         []Document          `json:"document_ids"`
	SelfSign            bool                `json:"self_sign"`
	DocumentFields      []DocumentField     `json:"document_fields"`
	TemplateIds         []string            `json:"template_ids"`
	InProcess           bool                `json:"in_process"`
	Validity            int                 `json:"validity"`
	RequestTypeName     string              `json:"request_type_name"`
	VisibleSignSettings VisibleSignSettings `json:"visible_sign_settings"`
	RequestId           string              `json:"request_id"`
	Zsdocumentid        string              `json:"zsdocumentid"`
	RequestTypeId       string              `json:"request_type_id"`
	OwnerLastName       string              `json:"owner_last_name"`
	Actions             []Action            `json:"actions"`
	AttachmentSize      int                 `json:"attachment_size"`
}

type Document struct {
	ImageString     string `json:"image_string"`
	DocumentName    string `json:"document_name"`
	Pages           []Page `json:"pages"`
	DocumentSize    int    `json:"document_size"`
	DocumentOrder   string `json:"document_order"`
	IsNom151Present bool   `json:"is_nom151_present"`
	IsEditable      bool   `json:"is_editable"`
	TotalPages      int    `json:"total_pages"`
	DocumentId      string `json:"document_id"`
}

type Page struct {
	ImageString string `json:"image_string"`
	Page        int    `json:"page"`
	IsThumbnail bool   `json:"is_thumbnail"`
}

type DocumentField struct {
	DocumentId string        `json:"document_id"`
	Fields     []interface{} `json:"fields"`
}

type VisibleSignSettings struct {
	AllowCustomSignerReasonVisibleSign bool     `json:"allow_custom_signer_reason_visible_sign"`
	AllowPredefinedReasonVisibleSign   bool     `json:"allow_predefined_reason_visible_sign"`
	VisibleSign                        bool     `json:"visible_sign"`
	PredefinedReasonsVisibleSign       []string `json:"predefined_reasons_visible_sign"`
	AllowReasonVisibleSign             bool     `json:"allow_reason_visible_sign"`
}

type Action struct {
	VerifyRecipient         bool    `json:"verify_recipient"`
	RecipientCountrycodeIso string  `json:"recipient_countrycode_iso"`
	ActionType              string  `json:"action_type"`
	PrivateNotes            string  `json:"private_notes"`
	CloudProviderName       string  `json:"cloud_provider_name"`
	HasPayment              bool    `json:"has_payment"`
	RecipientEmail          string  `json:"recipient_email"`
	SendCompletedDocument   bool    `json:"send_completed_document"`
	AllowSigning            bool    `json:"allow_signing"`
	RecipientPhonenumber    string  `json:"recipient_phonenumber"`
	IsBulk                  bool    `json:"is_bulk"`
	ActionId                string  `json:"action_id"`
	IsRevoked               bool    `json:"is_revoked"`
	IsEmbedded              bool    `json:"is_embedded"`
	SigningOrder            int     `json:"signing_order"`
	CloudProviderId         int     `json:"cloud_provider_id"`
	RecipientName           string  `json:"recipient_name"`
	Fields                  []Field `json:"fields"`
	DeliveryMode            string  `json:"delivery_mode"`
	ActionStatus            string  `json:"action_status"`
	IsSigningGroup          bool    `json:"is_signing_group"`
	RecipientCountrycode    string  `json:"recipient_countrycode"`
}

type Field struct {
	FieldId            string        `json:"field_id"`
	XCoord             int           `json:"x_coord"`
	FieldTypeId        string        `json:"field_type_id"`
	AbsHeight          int           `json:"abs_height"`
	FieldCategory      string        `json:"field_category"`
	IsDisabled         bool          `json:"is_disabled"`
	FieldLabel         string        `json:"field_label"`
	IsHidden           bool          `json:"is_hidden"`
	IsMandatory        bool          `json:"is_mandatory"`
	PageNo             int           `json:"page_no"`
	DocumentId         string        `json:"document_id"`
	FieldName          string        `json:"field_name"`
	YValue             float64       `json:"y_value"`
	AbsWidth           int           `json:"abs_width"`
	ActionId           string        `json:"action_id"`
	Width              float64       `json:"width"`
	YCoord             int           `json:"y_coord"`
	FieldTypeName      string        `json:"field_type_name"`
	DescriptionTooltip string        `json:"description_tooltip"`
	XValue             float64       `json:"x_value"`
	Height             float64       `json:"height"`
	TextProperty       *TextProperty `json:"text_property,omitempty"`
	DefaultValue       string        `json:"default_value,omitempty"`
}

type TextProperty struct {
	IsItalic       bool   `json:"is_italic"`
	MaxFieldLength int    `json:"max_field_length"`
	IsUnderline    bool   `json:"is_underline"`
	FontColor      string `json:"font_color"`
	IsFixedWidth   bool   `json:"is_fixed_width"`
	FontSize       int    `json:"font_size"`
	IsFixedHeight  bool   `json:"is_fixed_height"`
	IsReadOnly     bool   `json:"is_read_only"`
	IsBold         bool   `json:"is_bold"`
	Font           string `json:"font"`
}

// SendTemplateSignatureRequest payload for creating a new signature request using templates
type SendTemplateSignatureRequest struct {
	Templates TemplateSignature `json:"templates"`
}

type TemplateSignature struct {
	FieldData struct {
		FieldTextData struct {
		} `json:"field_text_data"`
		FieldBooleanData struct {
		} `json:"field_boolean_data"`
		FieldDateData struct {
		} `json:"field_date_data"`
		FieldRadioData struct {
		} `json:"field_radio_data"`
		FieldCheckboxgroupData struct {
		} `json:"field_checkboxgroup_data"`
	} `json:"field_data"`
	Notes   string           `json:"notes"`
	Actions []TemplateAction `json:"actions"`
}

type TemplateAction struct {
	RecipientName        string `json:"recipient_name"`
	RecipientEmail       string `json:"recipient_email"`
	ActionId             string `json:"action_id"`
	RecipientPhonenumber string `json:"recipient_phonenumber,omitempty"`
	RecipientCountrycode string `json:"recipient_countrycode,omitempty"`
	SigningOrder         int    `json:"signing_order"`
	Role                 string `json:"role"`
	VerifyRecipient      bool   `json:"verify_recipient"`
	PrivateNotes         string `json:"private_notes"`
	IsEmbedded           bool   `json:"is_embedded"`
}

type TemplateSignatureResponse struct {
	Code     int `json:"code"`
	Requests struct {
		RequestStatus     string   `json:"request_status"`
		Notes             string   `json:"notes"`
		ReminderPeriod    int      `json:"reminder_period"`
		OwnerId           string   `json:"owner_id"`
		Description       string   `json:"description"`
		RequestName       string   `json:"request_name"`
		ModifiedTime      int64    `json:"modified_time"`
		ActionTime        int64    `json:"action_time"`
		IsDeleted         bool     `json:"is_deleted"`
		ExpirationDays    int      `json:"expiration_days"`
		IsSequential      bool     `json:"is_sequential"`
		SignSubmittedTime int64    `json:"sign_submitted_time"`
		TemplatesUsed     []string `json:"templates_used"`
		OwnerFirstName    string   `json:"owner_first_name"`
		SignPercentage    float64  `json:"sign_percentage"`
		ExpireBy          int64    `json:"expire_by"`
		IsExpiring        bool     `json:"is_expiring"`
		OwnerEmail        string   `json:"owner_email"`
		CreatedTime       int64    `json:"created_time"`
		EmailReminders    bool     `json:"email_reminders"`
		DocumentIds       []struct {
			ImageString  string `json:"image_string"`
			DocumentName string `json:"document_name"`
			Pages        []struct {
				ImageString string `json:"image_string"`
				Page        int    `json:"page"`
				IsThumbnail bool   `json:"is_thumbnail"`
			} `json:"pages"`
			DocumentSize    int    `json:"document_size"`
			DocumentOrder   string `json:"document_order"`
			IsNom151Present bool   `json:"is_nom151_present"`
			IsEditable      bool   `json:"is_editable"`
			TotalPages      int    `json:"total_pages"`
			DocumentId      string `json:"document_id"`
		} `json:"document_ids"`
		SelfSign       bool `json:"self_sign"`
		DocumentFields []struct {
			DocumentId string        `json:"document_id"`
			Fields     []interface{} `json:"fields"`
		} `json:"document_fields"`
		TemplateIds         []string `json:"template_ids"`
		InProcess           bool     `json:"in_process"`
		Validity            int      `json:"validity"`
		RequestTypeName     string   `json:"request_type_name"`
		VisibleSignSettings struct {
			AllowCustomSignerReasonVisibleSign bool     `json:"allow_custom_signer_reason_visible_sign"`
			AllowPredefinedReasonVisibleSign   bool     `json:"allow_predefined_reason_visible_sign"`
			VisibleSign                        bool     `json:"visible_sign"`
			PredefinedReasonsVisibleSign       []string `json:"predefined_reasons_visible_sign"`
			AllowReasonVisibleSign             bool     `json:"allow_reason_visible_sign"`
		} `json:"visible_sign_settings"`
		RequestId     string `json:"request_id"`
		Zsdocumentid  string `json:"zsdocumentid"`
		RequestTypeId string `json:"request_type_id"`
		OwnerLastName string `json:"owner_last_name"`
		Actions       []struct {
			VerifyRecipient         bool   `json:"verify_recipient"`
			RecipientCountrycodeIso string `json:"recipient_countrycode_iso"`
			ActionType              string `json:"action_type"`
			PrivateNotes            string `json:"private_notes"`
			CloudProviderName       string `json:"cloud_provider_name"`
			HasPayment              bool   `json:"has_payment"`
			RecipientEmail          string `json:"recipient_email"`
			SendCompletedDocument   bool   `json:"send_completed_document"`
			AllowSigning            bool   `json:"allow_signing"`
			RecipientPhonenumber    string `json:"recipient_phonenumber"`
			IsBulk                  bool   `json:"is_bulk"`
			ActionId                string `json:"action_id"`
			IsRevoked               bool   `json:"is_revoked"`
			IsEmbedded              bool   `json:"is_embedded"`
			SigningOrder            int    `json:"signing_order"`
			CloudProviderId         int    `json:"cloud_provider_id"`
			RecipientName           string `json:"recipient_name"`
			Fields                  []struct {
				FieldId            string  `json:"field_id"`
				XCoord             int     `json:"x_coord"`
				FieldTypeId        string  `json:"field_type_id"`
				AbsHeight          int     `json:"abs_height"`
				FieldCategory      string  `json:"field_category"`
				FieldLabel         string  `json:"field_label"`
				IsMandatory        bool    `json:"is_mandatory"`
				PageNo             int     `json:"page_no"`
				DocumentId         string  `json:"document_id"`
				IsDraggable        bool    `json:"is_draggable,omitempty"`
				FieldName          string  `json:"field_name"`
				YValue             float64 `json:"y_value"`
				AbsWidth           int     `json:"abs_width"`
				ActionId           string  `json:"action_id"`
				Width              float64 `json:"width"`
				YCoord             int     `json:"y_coord"`
				FieldTypeName      string  `json:"field_type_name"`
				DescriptionTooltip string  `json:"description_tooltip"`
				XValue             float64 `json:"x_value"`
				IsResizable        bool    `json:"is_resizable,omitempty"`
				Height             float64 `json:"height"`
				TextProperty       struct {
					IsItalic       bool   `json:"is_italic"`
					MaxFieldLength int    `json:"max_field_length"`
					IsUnderline    bool   `json:"is_underline"`
					FontColor      string `json:"font_color"`
					IsFixedWidth   bool   `json:"is_fixed_width"`
					FontSize       int    `json:"font_size"`
					IsFixedHeight  bool   `json:"is_fixed_height"`
					IsReadOnly     bool   `json:"is_read_only"`
					IsBold         bool   `json:"is_bold"`
					Font           string `json:"font"`
				} `json:"text_property,omitempty"`
				NameFormat     string `json:"name_format,omitempty"`
				DefaultValue   string `json:"default_value,omitempty"`
				TimeZoneOffset int    `json:"time_zone_offset,omitempty"`
				TimeZone       string `json:"time_zone,omitempty"`
				FieldValue     string `json:"field_value,omitempty"`
				DateFormat     string `json:"date_format,omitempty"`
			} `json:"fields"`
			DeliveryMode         string `json:"delivery_mode"`
			ActionStatus         string `json:"action_status"`
			IsSigningGroup       bool   `json:"is_signing_group"`
			RecipientCountrycode string `json:"recipient_countrycode"`
		} `json:"actions"`
	} `json:"requests"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type EmbeddedSignedDocumentResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	SignUrl string `json:"sign_url"`
	Status  string `json:"status"`
}

type AuthRefreshResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	ApiDomain   string `json:"api_domain"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}
