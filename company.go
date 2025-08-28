// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dataleonlabs

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/dataleonlabs/dataleonlabs-go/internal/apijson"
	"github.com/dataleonlabs/dataleonlabs-go/internal/apiquery"
	"github.com/dataleonlabs/dataleonlabs-go/internal/requestconfig"
	"github.com/dataleonlabs/dataleonlabs-go/option"
	"github.com/dataleonlabs/dataleonlabs-go/packages/param"
	"github.com/dataleonlabs/dataleonlabs-go/packages/respjson"
)

// CompanyService contains methods and other services that help with interacting
// with the dataleonlabs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCompanyService] method instead.
type CompanyService struct {
	Options   []option.RequestOption
	Documents CompanyDocumentService
}

// NewCompanyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCompanyService(opts ...option.RequestOption) (r CompanyService) {
	r = CompanyService{}
	r.Options = opts
	r.Documents = NewCompanyDocumentService(opts...)
	return
}

// Create a new company
func (r *CompanyService) New(ctx context.Context, body CompanyNewParams, opts ...option.RequestOption) (res *Company, err error) {
	opts = append(r.Options[:], opts...)
	path := "companies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a company by ID
func (r *CompanyService) Get(ctx context.Context, companyID string, query CompanyGetParams, opts ...option.RequestOption) (res *Company, err error) {
	opts = append(r.Options[:], opts...)
	if companyID == "" {
		err = errors.New("missing required company_id parameter")
		return
	}
	path := fmt.Sprintf("companies/%s", companyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Update a company by ID
func (r *CompanyService) Update(ctx context.Context, companyID string, body CompanyUpdateParams, opts ...option.RequestOption) (res *Company, err error) {
	opts = append(r.Options[:], opts...)
	if companyID == "" {
		err = errors.New("missing required company_id parameter")
		return
	}
	path := fmt.Sprintf("companies/%s", companyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Get all companies
func (r *CompanyService) List(ctx context.Context, query CompanyListParams, opts ...option.RequestOption) (res *[]Company, err error) {
	opts = append(r.Options[:], opts...)
	path := "companies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a company by ID
func (r *CompanyService) Delete(ctx context.Context, companyID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if companyID == "" {
		err = errors.New("missing required company_id parameter")
		return
	}
	path := fmt.Sprintf("companies/%s", companyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Represents a record of suspicion raised during Anti-Money Laundering (AML)
// screening. Includes metadata such as risk score, origin, and linked watchlist
// types.
type AmlSuspicion struct {
	// Human-readable description or title for the suspicious finding.
	Caption string `json:"caption"`
	// Country associated with the suspicion (ISO 3166-1 alpha-2 code).
	Country string `json:"country"`
	// Gender associated with the suspicion, if applicable.
	Gender string `json:"gender"`
	// Nature of the relationship between the entity and the suspicious activity (e.g.,
	// "linked", "associated").
	Relation string `json:"relation"`
	// Version of the evaluation schema or rule engine used.
	Schema string `json:"schema"`
	// Risk score between 0.0 and 1 indicating the severity of the suspicion.
	Score float64 `json:"score"`
	// Source system or service providing this suspicion.
	Source string `json:"source"`
	// Status of the suspicion review process. Possible values: "true_positive",
	// "false_positive", "pending".
	//
	// Any of "true_positive", "false_positive", "pending".
	Status AmlSuspicionStatus `json:"status"`
	// Category of the suspicion. Possible values: "crime", "sanction", "pep",
	// "adverse_news", "other".
	//
	// Any of "crime", "sanction", "pep", "adverse_news", "other".
	Type AmlSuspicionType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caption     respjson.Field
		Country     respjson.Field
		Gender      respjson.Field
		Relation    respjson.Field
		Schema      respjson.Field
		Score       respjson.Field
		Source      respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AmlSuspicion) RawJSON() string { return r.JSON.raw }
func (r *AmlSuspicion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the suspicion review process. Possible values: "true_positive",
// "false_positive", "pending".
type AmlSuspicionStatus string

const (
	AmlSuspicionStatusTruePositive  AmlSuspicionStatus = "true_positive"
	AmlSuspicionStatusFalsePositive AmlSuspicionStatus = "false_positive"
	AmlSuspicionStatusPending       AmlSuspicionStatus = "pending"
)

// Category of the suspicion. Possible values: "crime", "sanction", "pep",
// "adverse_news", "other".
type AmlSuspicionType string

const (
	AmlSuspicionTypeCrime       AmlSuspicionType = "crime"
	AmlSuspicionTypeSanction    AmlSuspicionType = "sanction"
	AmlSuspicionTypePep         AmlSuspicionType = "pep"
	AmlSuspicionTypeAdverseNews AmlSuspicionType = "adverse_news"
	AmlSuspicionTypeOther       AmlSuspicionType = "other"
)

// Represents a certificate file associated with an individual or company.
type Certificat struct {
	// Unique identifier for the certificate.
	ID string `json:"id"`
	// Timestamp when the certificate was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Name of the certificate file.
	Filename string `json:"filename"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Filename    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Certificat) RawJSON() string { return r.JSON.raw }
func (r *Certificat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a verification check result.
type Check struct {
	// Indicates whether the result or data is masked/hidden.
	Masked bool `json:"masked"`
	// Additional message or explanation about the check result.
	Message string `json:"message"`
	// Name or type of the check performed.
	Name string `json:"name"`
	// Result of the check, true if passed.
	Validate bool `json:"validate"`
	// Importance or weight of the check, often used in scoring.
	Weight int64 `json:"weight"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Masked      respjson.Field
		Message     respjson.Field
		Name        respjson.Field
		Validate    respjson.Field
		Weight      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Check) RawJSON() string { return r.JSON.raw }
func (r *Check) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Company struct {
	// List of AML (Anti-Money Laundering) suspicion entries linked to the company,
	// including their details.
	AmlSuspicions []AmlSuspicion `json:"aml_suspicions"`
	// Digital certificate associated with the company, if any, including its creation
	// timestamp and filename.
	Certificat Certificat `json:"certificat"`
	// List of verification or validation checks applied to the company, including
	// their results and messages.
	Checks []Check `json:"checks"`
	// Main information about the company being registered, including legal name,
	// registration ID, and address.
	Company CompanyCompany `json:"company"`
	// All documents submitted or associated with the company, including their metadata
	// and processing status.
	Documents []GenericDocument `json:"documents"`
	// List of members or actors associated with the company, including personal and
	// ownership information.
	Members []CompanyMember `json:"members"`
	// Admin or internal portal URL for viewing the company's details, typically used
	// by internal users.
	PortalURL string `json:"portal_url"`
	// Custom key-value metadata fields associated with the company, allowing for
	// flexible data storage.
	Properties []Property `json:"properties"`
	// Risk assessment associated with the company, including a risk code, reason, and
	// confidence score.
	Risk Risk `json:"risk"`
	// Optional identifier indicating the source of the company record, useful for
	// tracking or integration purposes.
	SourceID string `json:"source_id"`
	// Technical metadata related to the request, such as IP address, QR code settings,
	// and callback URLs.
	TechnicalData TechnicalData `json:"technical_data"`
	// Public-facing webview URL for the company’s identification process, allowing
	// external access to the company data.
	WebviewURL string `json:"webview_url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AmlSuspicions respjson.Field
		Certificat    respjson.Field
		Checks        respjson.Field
		Company       respjson.Field
		Documents     respjson.Field
		Members       respjson.Field
		PortalURL     respjson.Field
		Properties    respjson.Field
		Risk          respjson.Field
		SourceID      respjson.Field
		TechnicalData respjson.Field
		WebviewURL    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Company) RawJSON() string { return r.JSON.raw }
func (r *Company) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Main information about the company being registered, including legal name,
// registration ID, and address.
type CompanyCompany struct {
	// Full registered address of the company.
	Address string `json:"address"`
	// Closure date of the company, if applicable.
	ClosureDate time.Time `json:"closure_date" format:"date"`
	// Trade or commercial name of the company.
	CommercialName string `json:"commercial_name"`
	// Contact information for the company, including email, phone number, and address.
	Contact CompanyCompanyContact `json:"contact"`
	// Country code where the company is registered.
	Country string `json:"country"`
	// Contact email address for the company.
	Email string `json:"email" format:"email"`
	// Number of employees in the company.
	Employees int64 `json:"employees"`
	// Employer Identification Number (EIN) or equivalent.
	EmployerIdentificationNumber string `json:"employer_identification_number"`
	// Indicates whether an insolvency procedure exists for the company.
	InsolvencyExists bool `json:"insolvency_exists"`
	// Indicates whether an insolvency procedure is ongoing for the company.
	InsolvencyOngoing bool `json:"insolvency_ongoing"`
	// Legal form or structure of the company (e.g., LLC, SARL).
	LegalForm string `json:"legal_form"`
	// Legal registered name of the company.
	Name string `json:"name"`
	// Contact phone number for the company, including country code.
	PhoneNumber string `json:"phone_number"`
	// Date when the company was officially registered.
	RegistrationDate time.Time `json:"registration_date" format:"date"`
	// Official company registration number or ID.
	RegistrationID string `json:"registration_id"`
	// Total share capital of the company, including currency.
	ShareCapital string `json:"share_capital"`
	// Current status of the company (e.g., active, inactive).
	Status string `json:"status"`
	// Tax identification number for the company.
	TaxIdentificationNumber string `json:"tax_identification_number"`
	// Type of company within the workspace, e.g., main or affiliated.
	Type string `json:"type"`
	// Official website URL of the company.
	WebsiteURL string `json:"website_url" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address                      respjson.Field
		ClosureDate                  respjson.Field
		CommercialName               respjson.Field
		Contact                      respjson.Field
		Country                      respjson.Field
		Email                        respjson.Field
		Employees                    respjson.Field
		EmployerIdentificationNumber respjson.Field
		InsolvencyExists             respjson.Field
		InsolvencyOngoing            respjson.Field
		LegalForm                    respjson.Field
		Name                         respjson.Field
		PhoneNumber                  respjson.Field
		RegistrationDate             respjson.Field
		RegistrationID               respjson.Field
		ShareCapital                 respjson.Field
		Status                       respjson.Field
		TaxIdentificationNumber      respjson.Field
		Type                         respjson.Field
		WebsiteURL                   respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyCompany) RawJSON() string { return r.JSON.raw }
func (r *CompanyCompany) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contact information for the company, including email, phone number, and address.
type CompanyCompanyContact struct {
	// Department of the contact person.
	Department string `json:"department"`
	// Email address of the contact person.
	Email string `json:"email" format:"email"`
	// First name of the contact person.
	FirstName string `json:"first_name"`
	// Last name of the contact person.
	LastName string `json:"last_name"`
	// Phone number of the contact person.
	PhoneNumber string `json:"phone_number"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Department  respjson.Field
		Email       respjson.Field
		FirstName   respjson.Field
		LastName    respjson.Field
		PhoneNumber respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyCompanyContact) RawJSON() string { return r.JSON.raw }
func (r *CompanyCompanyContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a member or actor of a company, including personal and ownership
// information.
type CompanyMember struct {
	ID string `json:"id" format:"uuid"`
	// Address of the member, which may include street, city, postal code, and country.
	Address string `json:"address"`
	// Birthday (available only if type = person)
	Birthday time.Time `json:"birthday" format:"date-time"`
	// Birthplace (available only if type = person)
	Birthplace string `json:"birthplace"`
	// ISO 3166-1 alpha-2 country code of the member's address (e.g., "FR" for France).
	Country string `json:"country"`
	// List of documents associated with the member, including their metadata and
	// processing status.
	Documents []GenericDocument `json:"documents"`
	// Email address of the member, which may be used for communication or verification
	// purposes.
	Email string `json:"email" format:"email"`
	// First name (available only if type = person)
	FirstName string `json:"first_name"`
	// Indicates whether the member is a beneficial owner of the company, meaning they
	// have significant control or ownership.
	IsBeneficialOwner bool `json:"is_beneficial_owner"`
	// Indicates whether the member is a delegator, meaning they have authority to act
	// on behalf of the company.
	IsDelegator bool `json:"is_delegator"`
	// Last name (available only if type = person)
	LastName string `json:"last_name"`
	// Indicates whether liveness verification was performed for the member, typically
	// in the context of identity checks.
	LivenessVerification bool `json:"liveness_verification"`
	// Company name (available only if type = company)
	Name string `json:"name"`
	// Percentage of ownership the member has in the company, expressed as an integer
	// between 0 and 100.
	OwnershipPercentage int64 `json:"ownership_percentage"`
	// Contact phone number of the member, including country code and area code.
	PhoneNumber string `json:"phone_number"`
	// Postal code of the member's address, typically a numeric or alphanumeric code.
	PostalCode string `json:"postal_code"`
	// Official registration identifier of the member, such as a national ID or company
	// registration number.
	RegistrationID string `json:"registration_id"`
	// Type of relationship the member has with the company, such as "shareholder",
	// "director", or "beneficial_owner".
	Relation string `json:"relation"`
	// Role of the member within the company, such as "legal_representative",
	// "director", or "manager".
	Roles string `json:"roles"`
	// Source of the data (e.g., government, user, company)
	//
	// Any of "gouve", "user", "company".
	Source string `json:"source"`
	// Current state of the member in the workflow, such as "WAITING", "STARTED",
	// "RUNNING", or "PROCESSED".
	State string `json:"state"`
	// Status of the member in the system, indicating whether they are approved,
	// pending, or rejected. Possible values include "approved", "need_review",
	// "rejected".
	Status string `json:"status"`
	// Member type (person or company)
	//
	// Any of "person", "company".
	Type string `json:"type"`
	// Identifier of the workspace to which the member belongs, used for organizational
	// purposes.
	WorkspaceID string `json:"workspace_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		Address              respjson.Field
		Birthday             respjson.Field
		Birthplace           respjson.Field
		Country              respjson.Field
		Documents            respjson.Field
		Email                respjson.Field
		FirstName            respjson.Field
		IsBeneficialOwner    respjson.Field
		IsDelegator          respjson.Field
		LastName             respjson.Field
		LivenessVerification respjson.Field
		Name                 respjson.Field
		OwnershipPercentage  respjson.Field
		PhoneNumber          respjson.Field
		PostalCode           respjson.Field
		RegistrationID       respjson.Field
		Relation             respjson.Field
		Roles                respjson.Field
		Source               respjson.Field
		State                respjson.Field
		Status               respjson.Field
		Type                 respjson.Field
		WorkspaceID          respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyMember) RawJSON() string { return r.JSON.raw }
func (r *CompanyMember) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a generic property key-value pair with a specified type.
type Property struct {
	// Name/key of the property.
	Name string `json:"name"`
	// Data type of the property value.
	Type string `json:"type"`
	// Value associated with the property name.
	Value string `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Property) RawJSON() string { return r.JSON.raw }
func (r *Property) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a risk assessment result, including a risk code, explanation, and a
// confidence score.
type Risk struct {
	// Risk category or code identifier.
	Code string `json:"code"`
	// Explanation or justification for the assigned risk.
	Reason string `json:"reason"`
	// Numeric risk score between 0.0 and 1.0 indicating severity or confidence.
	Score float64 `json:"score"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Reason      respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Risk) RawJSON() string { return r.JSON.raw }
func (r *Risk) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contains technical metadata related to processing and communication of an
// entity.
type TechnicalData struct {
	// Flag indicating whether there are active research AML (Anti-Money Laundering)
	// suspicions for the object when you apply for a new entry or get an existing one.
	ActiveAmlSuspicions bool `json:"active_aml_suspicions"`
	// Version number of the API used.
	APIVersion int64 `json:"api_version"`
	// Timestamp when the request or process was approved.
	ApprovedAt time.Time `json:"approved_at" format:"date-time"`
	// URL to receive callback data from the AML system.
	CallbackURL string `json:"callback_url" format:"uri"`
	// URL to receive notification updates about the processing status.
	CallbackURLNotification string `json:"callback_url_notification" format:"uri"`
	// Flag to indicate if notifications are disabled.
	DisableNotification bool `json:"disable_notification"`
	// Timestamp when notifications were disabled; null if never disabled.
	DisableNotificationDate time.Time `json:"disable_notification_date,nullable" format:"date-time"`
	// Export format defined by the API (e.g., "json", "xml").
	ExportType string `json:"export_type"`
	// Timestamp when the process finished.
	FinishedAt time.Time `json:"finished_at" format:"date-time"`
	// IP address of the our system handling the request.
	IP string `json:"ip"`
	// Language preference used in the client workspace (e.g., "fra").
	Language string `json:"language"`
	// IP address of the end client (final user) captured.
	LocationIP string `json:"location_ip"`
	// Timestamp indicating when the request or process needs review; null if none.
	NeedReviewAt time.Time `json:"need_review_at,nullable" format:"date-time"`
	// Flag indicating if notification confirmation is required or received.
	NotificationConfirmation bool `json:"notification_confirmation"`
	// Indicates whether QR code is enabled ("true" or "false").
	QrCode string `json:"qr_code"`
	// Flag indicating whether to include raw data in the response.
	RawData bool `json:"raw_data"`
	// Timestamp when the request or process was rejected; null if not rejected.
	RejectedAt time.Time `json:"rejected_at,nullable" format:"date-time"`
	// Duration of the user session in seconds.
	SessionDuration int64 `json:"session_duration"`
	// Timestamp when the process started.
	StartedAt time.Time `json:"started_at" format:"date-time"`
	// Date/time of data transfer.
	TransferAt time.Time `json:"transfer_at" format:"date-time"`
	// Mode of data transfer.
	TransferMode string `json:"transfer_mode"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActiveAmlSuspicions      respjson.Field
		APIVersion               respjson.Field
		ApprovedAt               respjson.Field
		CallbackURL              respjson.Field
		CallbackURLNotification  respjson.Field
		DisableNotification      respjson.Field
		DisableNotificationDate  respjson.Field
		ExportType               respjson.Field
		FinishedAt               respjson.Field
		IP                       respjson.Field
		Language                 respjson.Field
		LocationIP               respjson.Field
		NeedReviewAt             respjson.Field
		NotificationConfirmation respjson.Field
		QrCode                   respjson.Field
		RawData                  respjson.Field
		RejectedAt               respjson.Field
		SessionDuration          respjson.Field
		StartedAt                respjson.Field
		TransferAt               respjson.Field
		TransferMode             respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TechnicalData) RawJSON() string { return r.JSON.raw }
func (r *TechnicalData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyNewParams struct {
	// Main information about the company being registered.
	Company CompanyNewParamsCompany `json:"company,omitzero,required"`
	// Unique identifier of the workspace in which the company is being created.
	WorkspaceID string `json:"workspace_id,required"`
	// Optional identifier to track the origin of the request or integration from your
	// system.
	SourceID param.Opt[string] `json:"source_id,omitzero"`
	// Technical metadata and callback configuration.
	TechnicalData CompanyNewParamsTechnicalData `json:"technical_data,omitzero"`
	paramObj
}

func (r CompanyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CompanyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CompanyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Main information about the company being registered.
//
// The property Name is required.
type CompanyNewParamsCompany struct {
	// Legal name of the company.
	Name string `json:"name,required"`
	// Registered address of the company.
	Address param.Opt[string] `json:"address,omitzero"`
	// Commercial or trade name of the company, if different from the legal name.
	CommercialName param.Opt[string] `json:"commercial_name,omitzero"`
	// ISO 3166-1 alpha-2 country code of company registration (e.g., "FR" for France).
	Country param.Opt[string] `json:"country,omitzero"`
	// Contact email address for the company.
	Email param.Opt[string] `json:"email,omitzero" format:"email"`
	// Employer Identification Number (EIN) or equivalent.
	EmployerIdentificationNumber param.Opt[string] `json:"employer_identification_number,omitzero"`
	// Legal structure of the company (e.g., SARL, SAS).
	LegalForm param.Opt[string] `json:"legal_form,omitzero"`
	// Contact phone number for the company.
	PhoneNumber param.Opt[string] `json:"phone_number,omitzero" format:"string"`
	// Date of official company registration in YYYY-MM-DD format.
	RegistrationDate param.Opt[string] `json:"registration_date,omitzero"`
	// Official company registration identifier.
	RegistrationID param.Opt[string] `json:"registration_id,omitzero"`
	// Declared share capital of the company, usually in euros.
	ShareCapital param.Opt[string] `json:"share_capital,omitzero"`
	// Current status of the company (e.g., active, inactive).
	Status param.Opt[string] `json:"status,omitzero"`
	// National tax identifier (e.g., VAT or TIN).
	TaxIdentificationNumber param.Opt[string] `json:"tax_identification_number,omitzero"`
	// Type of company, such as "main" or "affiliated".
	Type param.Opt[string] `json:"type,omitzero"`
	// Company’s official website URL.
	WebsiteURL param.Opt[string] `json:"website_url,omitzero"`
	paramObj
}

func (r CompanyNewParamsCompany) MarshalJSON() (data []byte, err error) {
	type shadow CompanyNewParamsCompany
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CompanyNewParamsCompany) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Technical metadata and callback configuration.
type CompanyNewParamsTechnicalData struct {
	// Flag indicating whether there are active research AML (Anti-Money Laundering)
	// suspicions for the company when you apply for a new entry or get an existing
	// one.
	ActiveAmlSuspicions param.Opt[bool] `json:"active_aml_suspicions,omitzero"`
	// URL to receive a callback once the company is processed.
	CallbackURL param.Opt[string] `json:"callback_url,omitzero" format:"uri"`
	// URL to receive notifications about the processing state and status.
	CallbackURLNotification param.Opt[string] `json:"callback_url_notification,omitzero" format:"uri"`
	// Preferred language for responses or notifications (e.g., "eng", "fra").
	Language param.Opt[string] `json:"language,omitzero"`
	// Flag indicating whether to include raw data in the response.
	RawData param.Opt[bool] `json:"raw_data,omitzero"`
	paramObj
}

func (r CompanyNewParamsTechnicalData) MarshalJSON() (data []byte, err error) {
	type shadow CompanyNewParamsTechnicalData
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CompanyNewParamsTechnicalData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyGetParams struct {
	// Include document signed url
	Document param.Opt[bool] `query:"document,omitzero" json:"-"`
	// Scope filter (id or scope)
	Scope param.Opt[string] `query:"scope,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CompanyGetParams]'s query parameters as `url.Values`.
func (r CompanyGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CompanyUpdateParams struct {
	// Main information about the company being registered.
	Company CompanyUpdateParamsCompany `json:"company,omitzero,required"`
	// Unique identifier of the workspace in which the company is being created.
	WorkspaceID string `json:"workspace_id,required"`
	// Optional identifier to track the origin of the request or integration from your
	// system.
	SourceID param.Opt[string] `json:"source_id,omitzero"`
	// Technical metadata and callback configuration.
	TechnicalData CompanyUpdateParamsTechnicalData `json:"technical_data,omitzero"`
	paramObj
}

func (r CompanyUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow CompanyUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CompanyUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Main information about the company being registered.
//
// The property Name is required.
type CompanyUpdateParamsCompany struct {
	// Legal name of the company.
	Name string `json:"name,required"`
	// Registered address of the company.
	Address param.Opt[string] `json:"address,omitzero"`
	// Commercial or trade name of the company, if different from the legal name.
	CommercialName param.Opt[string] `json:"commercial_name,omitzero"`
	// ISO 3166-1 alpha-2 country code of company registration (e.g., "FR" for France).
	Country param.Opt[string] `json:"country,omitzero"`
	// Contact email address for the company.
	Email param.Opt[string] `json:"email,omitzero" format:"email"`
	// Employer Identification Number (EIN) or equivalent.
	EmployerIdentificationNumber param.Opt[string] `json:"employer_identification_number,omitzero"`
	// Legal structure of the company (e.g., SARL, SAS).
	LegalForm param.Opt[string] `json:"legal_form,omitzero"`
	// Contact phone number for the company.
	PhoneNumber param.Opt[string] `json:"phone_number,omitzero" format:"string"`
	// Date of official company registration in YYYY-MM-DD format.
	RegistrationDate param.Opt[string] `json:"registration_date,omitzero"`
	// Official company registration identifier.
	RegistrationID param.Opt[string] `json:"registration_id,omitzero"`
	// Declared share capital of the company, usually in euros.
	ShareCapital param.Opt[string] `json:"share_capital,omitzero"`
	// Current status of the company (e.g., active, inactive).
	Status param.Opt[string] `json:"status,omitzero"`
	// National tax identifier (e.g., VAT or TIN).
	TaxIdentificationNumber param.Opt[string] `json:"tax_identification_number,omitzero"`
	// Type of company, such as "main" or "affiliated".
	Type param.Opt[string] `json:"type,omitzero"`
	// Company’s official website URL.
	WebsiteURL param.Opt[string] `json:"website_url,omitzero"`
	paramObj
}

func (r CompanyUpdateParamsCompany) MarshalJSON() (data []byte, err error) {
	type shadow CompanyUpdateParamsCompany
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CompanyUpdateParamsCompany) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Technical metadata and callback configuration.
type CompanyUpdateParamsTechnicalData struct {
	// Flag indicating whether there are active research AML (Anti-Money Laundering)
	// suspicions for the company when you apply for a new entry or get an existing
	// one.
	ActiveAmlSuspicions param.Opt[bool] `json:"active_aml_suspicions,omitzero"`
	// URL to receive a callback once the company is processed.
	CallbackURL param.Opt[string] `json:"callback_url,omitzero" format:"uri"`
	// URL to receive notifications about the processing state and status.
	CallbackURLNotification param.Opt[string] `json:"callback_url_notification,omitzero" format:"uri"`
	// Preferred language for responses or notifications (e.g., "eng", "fra").
	Language param.Opt[string] `json:"language,omitzero"`
	// Flag indicating whether to include raw data in the response.
	RawData param.Opt[bool] `json:"raw_data,omitzero"`
	paramObj
}

func (r CompanyUpdateParamsTechnicalData) MarshalJSON() (data []byte, err error) {
	type shadow CompanyUpdateParamsTechnicalData
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CompanyUpdateParamsTechnicalData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyListParams struct {
	// Filter companies created before this date (format YYYY-MM-DD)
	EndDate param.Opt[time.Time] `query:"end_date,omitzero" format:"date" json:"-"`
	// Number of results to return (between 1 and 100)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of results to skip (must be ≥ 0)
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Filter by source ID
	SourceID param.Opt[string] `query:"source_id,omitzero" json:"-"`
	// Filter companies created after this date (format YYYY-MM-DD)
	StartDate param.Opt[time.Time] `query:"start_date,omitzero" format:"date" json:"-"`
	// Filter by workspace ID
	WorkspaceID param.Opt[string] `query:"workspace_id,omitzero" json:"-"`
	// Filter by company state (must be one of the allowed values)
	//
	// Any of "VOID", "WAITING", "STARTED", "RUNNING", "PROCESSED", "FAILED",
	// "ABORTED", "EXPIRED", "DELETED".
	State CompanyListParamsState `query:"state,omitzero" json:"-"`
	// Filter by individual status (must be one of the allowed values)
	//
	// Any of "rejected", "need_review", "approved".
	Status CompanyListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CompanyListParams]'s query parameters as `url.Values`.
func (r CompanyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by company state (must be one of the allowed values)
type CompanyListParamsState string

const (
	CompanyListParamsStateVoid      CompanyListParamsState = "VOID"
	CompanyListParamsStateWaiting   CompanyListParamsState = "WAITING"
	CompanyListParamsStateStarted   CompanyListParamsState = "STARTED"
	CompanyListParamsStateRunning   CompanyListParamsState = "RUNNING"
	CompanyListParamsStateProcessed CompanyListParamsState = "PROCESSED"
	CompanyListParamsStateFailed    CompanyListParamsState = "FAILED"
	CompanyListParamsStateAborted   CompanyListParamsState = "ABORTED"
	CompanyListParamsStateExpired   CompanyListParamsState = "EXPIRED"
	CompanyListParamsStateDeleted   CompanyListParamsState = "DELETED"
)

// Filter by individual status (must be one of the allowed values)
type CompanyListParamsStatus string

const (
	CompanyListParamsStatusRejected   CompanyListParamsStatus = "rejected"
	CompanyListParamsStatusNeedReview CompanyListParamsStatus = "need_review"
	CompanyListParamsStatusApproved   CompanyListParamsStatus = "approved"
)
