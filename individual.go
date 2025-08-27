// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dataleonlabs

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/dataleonlabs-go/internal/apijson"
	"github.com/stainless-sdks/dataleonlabs-go/internal/apiquery"
	"github.com/stainless-sdks/dataleonlabs-go/internal/requestconfig"
	"github.com/stainless-sdks/dataleonlabs-go/option"
	"github.com/stainless-sdks/dataleonlabs-go/packages/param"
	"github.com/stainless-sdks/dataleonlabs-go/packages/respjson"
)

// IndividualService contains methods and other services that help with interacting
// with the dataleonlabs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIndividualService] method instead.
type IndividualService struct {
	Options   []option.RequestOption
	Documents IndividualDocumentService
}

// NewIndividualService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIndividualService(opts ...option.RequestOption) (r IndividualService) {
	r = IndividualService{}
	r.Options = opts
	r.Documents = NewIndividualDocumentService(opts...)
	return
}

// Create a new individual
func (r *IndividualService) New(ctx context.Context, body IndividualNewParams, opts ...option.RequestOption) (res *Individual, err error) {
	opts = append(r.Options[:], opts...)
	path := "individuals"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get an individual by ID
func (r *IndividualService) Get(ctx context.Context, individualID string, query IndividualGetParams, opts ...option.RequestOption) (res *Individual, err error) {
	opts = append(r.Options[:], opts...)
	if individualID == "" {
		err = errors.New("missing required individual_id parameter")
		return
	}
	path := fmt.Sprintf("individuals/%s", individualID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Update an individual by ID
func (r *IndividualService) Update(ctx context.Context, individualID string, body IndividualUpdateParams, opts ...option.RequestOption) (res *Individual, err error) {
	opts = append(r.Options[:], opts...)
	if individualID == "" {
		err = errors.New("missing required individual_id parameter")
		return
	}
	path := fmt.Sprintf("individuals/%s", individualID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Get all individuals
func (r *IndividualService) List(ctx context.Context, query IndividualListParams, opts ...option.RequestOption) (res *[]Individual, err error) {
	opts = append(r.Options[:], opts...)
	path := "individuals"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete an individual by ID
func (r *IndividualService) Delete(ctx context.Context, individualID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if individualID == "" {
		err = errors.New("missing required individual_id parameter")
		return
	}
	path := fmt.Sprintf("individuals/%s", individualID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Represents a single individual record, including identification, status, and
// associated metadata.
type Individual struct {
	// Unique identifier of the individual.
	ID string `json:"id" format:"uuid"`
	// List of AML (Anti-Money Laundering) suspicion entries linked to the individual.
	AmlSuspicions []AmlSuspicion `json:"aml_suspicions"`
	// URL to authenticate the individual, usually for document signing or onboarding.
	AuthURL string `json:"auth_url" format:"uri"`
	// Digital certificate associated with the individual, if any.
	Certificat Certificat `json:"certificat"`
	// List of verification or validation checks applied to the individual.
	Checks []Check `json:"checks"`
	// Timestamp of the individual's creation in ISO 8601 format.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// All documents submitted or associated with the individual.
	Documents []GenericDocument `json:"documents"`
	// Reference to the individual's identity document.
	IdentityCard IndividualIdentityCard `json:"identity_card"`
	// Internal sequential number or reference for the individual.
	Number int64 `json:"number"`
	// Personal details of the individual, such as name, date of birth, and contact
	// info.
	Person IndividualPerson `json:"person"`
	// Admin or internal portal URL for viewing the individual's details.
	PortalURL string `json:"portal_url" format:"uri"`
	// Custom key-value metadata fields associated with the individual.
	Properties []Property `json:"properties"`
	// Risk assessment associated with the individual.
	Risk Risk `json:"risk"`
	// Optional identifier indicating the source of the individual record.
	SourceID string `json:"source_id"`
	// Current operational state in the workflow (e.g., WAITING, IN_PROGRESS,
	// COMPLETED).
	State string `json:"state"`
	// Overall processing status of the individual (e.g., rejected, need_review,
	// approved).
	Status string `json:"status"`
	// List of tags assigned to the individual for categorization or metadata purposes.
	Tags []IndividualTag `json:"tags"`
	// Technical metadata related to the request (e.g., QR code settings, language).
	TechnicalData TechnicalData `json:"technical_data"`
	// Public-facing webview URL for the individual’s identification process.
	WebviewURL string `json:"webview_url" format:"uri"`
	// Identifier of the workspace to which the individual belongs.
	WorkspaceID string `json:"workspace_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		AmlSuspicions respjson.Field
		AuthURL       respjson.Field
		Certificat    respjson.Field
		Checks        respjson.Field
		CreatedAt     respjson.Field
		Documents     respjson.Field
		IdentityCard  respjson.Field
		Number        respjson.Field
		Person        respjson.Field
		PortalURL     respjson.Field
		Properties    respjson.Field
		Risk          respjson.Field
		SourceID      respjson.Field
		State         respjson.Field
		Status        respjson.Field
		Tags          respjson.Field
		TechnicalData respjson.Field
		WebviewURL    respjson.Field
		WorkspaceID   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Individual) RawJSON() string { return r.JSON.raw }
func (r *Individual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Reference to the individual's identity document.
type IndividualIdentityCard struct {
	// Unique identifier for the document.
	ID string `json:"id"`
	// Signed URL linking to the back image of the document.
	BackDocumentSignedURL string `json:"back_document_signed_url" format:"uri"`
	// Place of birth as indicated on the document.
	BirthPlace string `json:"birth_place"`
	// Date of birth in DD/MM/YYYY format as shown on the document.
	Birthday string `json:"birthday"`
	// Country code issuing the document (ISO 3166-1 alpha-2).
	Country string `json:"country"`
	// Expiration date of the document, in YYYY-MM-DD format.
	ExpirationDate string `json:"expiration_date"`
	// First name as shown on the document.
	FirstName string `json:"first_name"`
	// Signed URL linking to the front image of the document.
	FrontDocumentSignedURL string `json:"front_document_signed_url" format:"uri"`
	// Gender indicated on the document (e.g., "M" or "F").
	Gender string `json:"gender"`
	// Date when the document was issued, in YYYY-MM-DD format.
	IssueDate string `json:"issue_date"`
	// Last name as shown on the document.
	LastName string `json:"last_name"`
	// First line of the Machine Readable Zone (MRZ) on the document.
	MrzLine1 string `json:"mrz_line_1"`
	// Second line of the MRZ on the document.
	MrzLine2 string `json:"mrz_line_2"`
	// Third line of the MRZ if applicable; otherwise null.
	MrzLine3 string `json:"mrz_line_3,nullable"`
	// Type of document (e.g., passport, identity card).
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		BackDocumentSignedURL  respjson.Field
		BirthPlace             respjson.Field
		Birthday               respjson.Field
		Country                respjson.Field
		ExpirationDate         respjson.Field
		FirstName              respjson.Field
		FrontDocumentSignedURL respjson.Field
		Gender                 respjson.Field
		IssueDate              respjson.Field
		LastName               respjson.Field
		MrzLine1               respjson.Field
		MrzLine2               respjson.Field
		MrzLine3               respjson.Field
		Type                   respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IndividualIdentityCard) RawJSON() string { return r.JSON.raw }
func (r *IndividualIdentityCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Personal details of the individual, such as name, date of birth, and contact
// info.
type IndividualPerson struct {
	// Date of birth, formatted as DD/MM/YYYY.
	Birthday string `json:"birthday"`
	// Email address of the individual.
	Email string `json:"email" format:"email"`
	// Signed URL linking to the person’s face image.
	FaceImageSignedURL string `json:"face_image_signed_url" format:"uri"`
	// First (given) name of the person.
	FirstName string `json:"first_name"`
	// Full name of the person, typically concatenation of first and last names.
	FullName string `json:"full_name"`
	// Gender of the individual (e.g., "M" for male, "F" for female).
	Gender string `json:"gender"`
	// Last (family) name of the person.
	LastName string `json:"last_name"`
	// Maiden name of the person, if applicable.
	MaidenName string `json:"maiden_name"`
	// Contact phone number including country code.
	PhoneNumber string `json:"phone_number"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Birthday           respjson.Field
		Email              respjson.Field
		FaceImageSignedURL respjson.Field
		FirstName          respjson.Field
		FullName           respjson.Field
		Gender             respjson.Field
		LastName           respjson.Field
		MaidenName         respjson.Field
		PhoneNumber        respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IndividualPerson) RawJSON() string { return r.JSON.raw }
func (r *IndividualPerson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a key-value metadata tag that can be associated with entities such as
// individuals or companies.
type IndividualTag struct {
	// Name of the tag used to identify the metadata field.
	Key string `json:"key"`
	// Indicates whether the tag is private (not visible to external users).
	Private bool `json:"private"`
	// Data type of the tag value (e.g., "string", "number", "boolean").
	Type string `json:"type"`
	// Value assigned to the tag.
	Value string `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Private     respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IndividualTag) RawJSON() string { return r.JSON.raw }
func (r *IndividualTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IndividualNewParams struct {
	// Unique identifier of the workspace where the individual is being registered.
	WorkspaceID string `json:"workspace_id,required"`
	// Optional identifier for tracking the source system or integration from your
	// system.
	SourceID param.Opt[string] `json:"source_id,omitzero"`
	// Personal information about the individual.
	Person IndividualNewParamsPerson `json:"person,omitzero"`
	// Technical metadata related to the request or processing.
	TechnicalData IndividualNewParamsTechnicalData `json:"technical_data,omitzero"`
	paramObj
}

func (r IndividualNewParams) MarshalJSON() (data []byte, err error) {
	type shadow IndividualNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IndividualNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Personal information about the individual.
type IndividualNewParamsPerson struct {
	// Date of birth in DD/MM/YYYY format.
	Birthday param.Opt[string] `json:"birthday,omitzero"`
	// Email address of the individual.
	Email param.Opt[string] `json:"email,omitzero" format:"email"`
	// First name of the individual.
	FirstName param.Opt[string] `json:"first_name,omitzero"`
	// Last name (family name) of the individual.
	LastName param.Opt[string] `json:"last_name,omitzero"`
	// Maiden name, if applicable.
	MaidenName param.Opt[string] `json:"maiden_name,omitzero"`
	// Phone number of the individual.
	PhoneNumber param.Opt[string] `json:"phone_number,omitzero"`
	// Gender of the individual (M for male, F for female).
	//
	// Any of "M", "F".
	Gender string `json:"gender,omitzero"`
	paramObj
}

func (r IndividualNewParamsPerson) MarshalJSON() (data []byte, err error) {
	type shadow IndividualNewParamsPerson
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IndividualNewParamsPerson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[IndividualNewParamsPerson](
		"gender", "M", "F",
	)
}

// Technical metadata related to the request or processing.
type IndividualNewParamsTechnicalData struct {
	// Flag indicating whether there are active research AML (Anti-Money Laundering)
	// suspicions for the individual when you apply for a new entry or get an existing
	// one.
	ActiveAmlSuspicions param.Opt[bool] `json:"active_aml_suspicions,omitzero"`
	// URL to call back upon completion of processing.
	CallbackURL param.Opt[string] `json:"callback_url,omitzero" format:"uri"`
	// URL for receive notifications about the processing state or status.
	CallbackURLNotification param.Opt[string] `json:"callback_url_notification,omitzero" format:"uri"`
	// Preferred language for communication (e.g., "eng", "fra").
	Language param.Opt[string] `json:"language,omitzero"`
	// Flag indicating whether to include raw data in the response.
	RawData param.Opt[bool] `json:"raw_data,omitzero"`
	paramObj
}

func (r IndividualNewParamsTechnicalData) MarshalJSON() (data []byte, err error) {
	type shadow IndividualNewParamsTechnicalData
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IndividualNewParamsTechnicalData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IndividualGetParams struct {
	// Include document information
	Document param.Opt[bool] `query:"document,omitzero" json:"-"`
	// Scope filter (id or scope)
	Scope param.Opt[string] `query:"scope,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IndividualGetParams]'s query parameters as `url.Values`.
func (r IndividualGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IndividualUpdateParams struct {
	// Unique identifier of the workspace where the individual is being registered.
	WorkspaceID string `json:"workspace_id,required"`
	// Optional identifier for tracking the source system or integration from your
	// system.
	SourceID param.Opt[string] `json:"source_id,omitzero"`
	// Personal information about the individual.
	Person IndividualUpdateParamsPerson `json:"person,omitzero"`
	// Technical metadata related to the request or processing.
	TechnicalData IndividualUpdateParamsTechnicalData `json:"technical_data,omitzero"`
	paramObj
}

func (r IndividualUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow IndividualUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IndividualUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Personal information about the individual.
type IndividualUpdateParamsPerson struct {
	// Date of birth in DD/MM/YYYY format.
	Birthday param.Opt[string] `json:"birthday,omitzero"`
	// Email address of the individual.
	Email param.Opt[string] `json:"email,omitzero" format:"email"`
	// First name of the individual.
	FirstName param.Opt[string] `json:"first_name,omitzero"`
	// Last name (family name) of the individual.
	LastName param.Opt[string] `json:"last_name,omitzero"`
	// Maiden name, if applicable.
	MaidenName param.Opt[string] `json:"maiden_name,omitzero"`
	// Phone number of the individual.
	PhoneNumber param.Opt[string] `json:"phone_number,omitzero"`
	// Gender of the individual (M for male, F for female).
	//
	// Any of "M", "F".
	Gender string `json:"gender,omitzero"`
	paramObj
}

func (r IndividualUpdateParamsPerson) MarshalJSON() (data []byte, err error) {
	type shadow IndividualUpdateParamsPerson
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IndividualUpdateParamsPerson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[IndividualUpdateParamsPerson](
		"gender", "M", "F",
	)
}

// Technical metadata related to the request or processing.
type IndividualUpdateParamsTechnicalData struct {
	// Flag indicating whether there are active research AML (Anti-Money Laundering)
	// suspicions for the individual when you apply for a new entry or get an existing
	// one.
	ActiveAmlSuspicions param.Opt[bool] `json:"active_aml_suspicions,omitzero"`
	// URL to call back upon completion of processing.
	CallbackURL param.Opt[string] `json:"callback_url,omitzero" format:"uri"`
	// URL for receive notifications about the processing state or status.
	CallbackURLNotification param.Opt[string] `json:"callback_url_notification,omitzero" format:"uri"`
	// Preferred language for communication (e.g., "eng", "fra").
	Language param.Opt[string] `json:"language,omitzero"`
	// Flag indicating whether to include raw data in the response.
	RawData param.Opt[bool] `json:"raw_data,omitzero"`
	paramObj
}

func (r IndividualUpdateParamsTechnicalData) MarshalJSON() (data []byte, err error) {
	type shadow IndividualUpdateParamsTechnicalData
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IndividualUpdateParamsTechnicalData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IndividualListParams struct {
	// Filter individuals created before this date (format YYYY-MM-DD)
	EndDate param.Opt[time.Time] `query:"end_date,omitzero" format:"date" json:"-"`
	// Number of results to return (between 1 and 100)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of results to offset (must be ≥ 0)
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Filter by source ID
	SourceID param.Opt[string] `query:"source_id,omitzero" json:"-"`
	// Filter individuals created after this date (format YYYY-MM-DD)
	StartDate param.Opt[time.Time] `query:"start_date,omitzero" format:"date" json:"-"`
	// Filter by workspace ID
	WorkspaceID param.Opt[string] `query:"workspace_id,omitzero" json:"-"`
	// Filter by individual status (must be one of the allowed values)
	//
	// Any of "VOID", "WAITING", "STARTED", "RUNNING", "PROCESSED", "FAILED",
	// "ABORTED", "EXPIRED", "DELETED".
	State IndividualListParamsState `query:"state,omitzero" json:"-"`
	// Filter by individual status (must be one of the allowed values)
	//
	// Any of "rejected", "need_review", "approved".
	Status IndividualListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IndividualListParams]'s query parameters as `url.Values`.
func (r IndividualListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by individual status (must be one of the allowed values)
type IndividualListParamsState string

const (
	IndividualListParamsStateVoid      IndividualListParamsState = "VOID"
	IndividualListParamsStateWaiting   IndividualListParamsState = "WAITING"
	IndividualListParamsStateStarted   IndividualListParamsState = "STARTED"
	IndividualListParamsStateRunning   IndividualListParamsState = "RUNNING"
	IndividualListParamsStateProcessed IndividualListParamsState = "PROCESSED"
	IndividualListParamsStateFailed    IndividualListParamsState = "FAILED"
	IndividualListParamsStateAborted   IndividualListParamsState = "ABORTED"
	IndividualListParamsStateExpired   IndividualListParamsState = "EXPIRED"
	IndividualListParamsStateDeleted   IndividualListParamsState = "DELETED"
)

// Filter by individual status (must be one of the allowed values)
type IndividualListParamsStatus string

const (
	IndividualListParamsStatusRejected   IndividualListParamsStatus = "rejected"
	IndividualListParamsStatusNeedReview IndividualListParamsStatus = "need_review"
	IndividualListParamsStatusApproved   IndividualListParamsStatus = "approved"
)
