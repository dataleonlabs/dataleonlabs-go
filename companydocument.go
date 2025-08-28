// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dataleonlabs

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/dataleonlabs/dataleonlabs-go/internal/apiform"
	"github.com/dataleonlabs/dataleonlabs-go/internal/apijson"
	"github.com/dataleonlabs/dataleonlabs-go/internal/requestconfig"
	"github.com/dataleonlabs/dataleonlabs-go/option"
	"github.com/dataleonlabs/dataleonlabs-go/packages/param"
	"github.com/dataleonlabs/dataleonlabs-go/packages/respjson"
)

// CompanyDocumentService contains methods and other services that help with
// interacting with the dataleonlabs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCompanyDocumentService] method instead.
type CompanyDocumentService struct {
	Options []option.RequestOption
}

// NewCompanyDocumentService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCompanyDocumentService(opts ...option.RequestOption) (r CompanyDocumentService) {
	r = CompanyDocumentService{}
	r.Options = opts
	return
}

// Get documents to an company
func (r *CompanyDocumentService) List(ctx context.Context, companyID string, opts ...option.RequestOption) (res *DocumentResponse, err error) {
	opts = append(r.Options[:], opts...)
	if companyID == "" {
		err = errors.New("missing required company_id parameter")
		return
	}
	path := fmt.Sprintf("companies/%s/documents", companyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Upload documents to an company
func (r *CompanyDocumentService) Upload(ctx context.Context, companyID string, body CompanyDocumentUploadParams, opts ...option.RequestOption) (res *GenericDocument, err error) {
	opts = append(r.Options[:], opts...)
	if companyID == "" {
		err = errors.New("missing required company_id parameter")
		return
	}
	path := fmt.Sprintf("companies/%s/documents", companyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type DocumentResponse struct {
	// List of documents associated with the response.
	Documents []DocumentResponseDocument `json:"documents"`
	// Total number of documents available in the response.
	TotalDocument int64 `json:"total_document"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Documents     respjson.Field
		TotalDocument respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentResponse) RawJSON() string { return r.JSON.raw }
func (r *DocumentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a document stored and processed by the system, such as an identity
// card or a PDF contract.
type DocumentResponseDocument struct {
	// Unique identifier of the document.
	ID string `json:"id"`
	// Functional type of the document (e.g., identity document, invoice).
	DocumentType string `json:"document_type"`
	// Original filename of the uploaded document.
	Filename string `json:"filename"`
	// Human-readable name of the document.
	Name string `json:"name"`
	// Secure URL to access the document.
	SignedURL string `json:"signed_url" format:"uri"`
	// Processing state of the document (e.g., WAITING, STARTED, RUNNING, PROCESSED).
	State string `json:"state"`
	// Validation status of the document (e.g., need_review, approved, rejected).
	Status string `json:"status"`
	// Identifier of the workspace to which the document belongs.
	WorkspaceID string `json:"workspace_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		DocumentType respjson.Field
		Filename     respjson.Field
		Name         respjson.Field
		SignedURL    respjson.Field
		State        respjson.Field
		Status       respjson.Field
		WorkspaceID  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentResponseDocument) RawJSON() string { return r.JSON.raw }
func (r *DocumentResponseDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a general document with metadata, verification checks, and extracted
// data.
type GenericDocument struct {
	// Unique identifier of the document.
	ID string `json:"id"`
	// List of verification checks performed on the document.
	Checks []Check `json:"checks"`
	// Timestamp when the document was created or uploaded.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Type/category of the document.
	DocumentType string `json:"document_type"`
	// Name or label for the document.
	Name string `json:"name"`
	// Signed URL for accessing the document file.
	SignedURL string `json:"signed_url" format:"uri"`
	// Current processing state of the document (e.g., WAITING, PROCESSED).
	State string `json:"state"`
	// Status of the document reception or approval.
	Status string `json:"status"`
	// List of tables extracted from the document, each containing operations.
	Tables []GenericDocumentTable `json:"tables"`
	// Extracted key-value pairs from the document, including confidence scores.
	Values []GenericDocumentValue `json:"values"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Checks       respjson.Field
		CreatedAt    respjson.Field
		DocumentType respjson.Field
		Name         respjson.Field
		SignedURL    respjson.Field
		State        respjson.Field
		Status       respjson.Field
		Tables       respjson.Field
		Values       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GenericDocument) RawJSON() string { return r.JSON.raw }
func (r *GenericDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GenericDocumentTable struct {
	// List of operations or actions associated with the table.
	Operation []any `json:"operation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operation   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GenericDocumentTable) RawJSON() string { return r.JSON.raw }
func (r *GenericDocumentTable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GenericDocumentValue struct {
	// Confidence score (between 0 and 1) for the extracted value.
	Confidence float64 `json:"confidence"`
	// Name or label of the extracted field.
	Name string `json:"name"`
	// List of integer values related to the field (e.g., bounding box coordinates).
	Value []int64 `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Confidence  respjson.Field
		Name        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GenericDocumentValue) RawJSON() string { return r.JSON.raw }
func (r *GenericDocumentValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyDocumentUploadParams struct {
	// Filter by document type for upload (must be one of the allowed values)
	//
	// Any of "liasse_fiscale", "amortised_loan_schedule", "invoice", "receipt",
	// "company_statuts", "registration_company_certificate", "kbis", "rib",
	// "livret_famille", "birth_certificate", "payslip", "social_security_card",
	// "vehicle_registration_certificate", "carte_grise", "criminal_record_extract",
	// "proof_of_address", "identity_card_front", "identity_card_back",
	// "driver_license_front", "driver_license_back", "identity_document",
	// "driver_license", "passport", "tax", "certificate_of_incorporation",
	// "certificate_of_good_standing", "lcb_ft_lab_aml_policies", "niu_entreprise",
	// "financial_statements", "rccm", "proof_of_source_funds", "organizational_chart",
	// "risk_policies".
	DocumentType CompanyDocumentUploadParamsDocumentType `json:"document_type,omitzero,required"`
	// URL of the file to upload (either `file` or `url` is required)
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	// File to upload (required)
	File io.Reader `json:"file,omitzero" format:"binary"`
	paramObj
}

func (r CompanyDocumentUploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// Filter by document type for upload (must be one of the allowed values)
type CompanyDocumentUploadParamsDocumentType string

const (
	CompanyDocumentUploadParamsDocumentTypeLiasseFiscale                  CompanyDocumentUploadParamsDocumentType = "liasse_fiscale"
	CompanyDocumentUploadParamsDocumentTypeAmortisedLoanSchedule          CompanyDocumentUploadParamsDocumentType = "amortised_loan_schedule"
	CompanyDocumentUploadParamsDocumentTypeInvoice                        CompanyDocumentUploadParamsDocumentType = "invoice"
	CompanyDocumentUploadParamsDocumentTypeReceipt                        CompanyDocumentUploadParamsDocumentType = "receipt"
	CompanyDocumentUploadParamsDocumentTypeCompanyStatuts                 CompanyDocumentUploadParamsDocumentType = "company_statuts"
	CompanyDocumentUploadParamsDocumentTypeRegistrationCompanyCertificate CompanyDocumentUploadParamsDocumentType = "registration_company_certificate"
	CompanyDocumentUploadParamsDocumentTypeKbis                           CompanyDocumentUploadParamsDocumentType = "kbis"
	CompanyDocumentUploadParamsDocumentTypeRib                            CompanyDocumentUploadParamsDocumentType = "rib"
	CompanyDocumentUploadParamsDocumentTypeLivretFamille                  CompanyDocumentUploadParamsDocumentType = "livret_famille"
	CompanyDocumentUploadParamsDocumentTypeBirthCertificate               CompanyDocumentUploadParamsDocumentType = "birth_certificate"
	CompanyDocumentUploadParamsDocumentTypePayslip                        CompanyDocumentUploadParamsDocumentType = "payslip"
	CompanyDocumentUploadParamsDocumentTypeSocialSecurityCard             CompanyDocumentUploadParamsDocumentType = "social_security_card"
	CompanyDocumentUploadParamsDocumentTypeVehicleRegistrationCertificate CompanyDocumentUploadParamsDocumentType = "vehicle_registration_certificate"
	CompanyDocumentUploadParamsDocumentTypeCarteGrise                     CompanyDocumentUploadParamsDocumentType = "carte_grise"
	CompanyDocumentUploadParamsDocumentTypeCriminalRecordExtract          CompanyDocumentUploadParamsDocumentType = "criminal_record_extract"
	CompanyDocumentUploadParamsDocumentTypeProofOfAddress                 CompanyDocumentUploadParamsDocumentType = "proof_of_address"
	CompanyDocumentUploadParamsDocumentTypeIdentityCardFront              CompanyDocumentUploadParamsDocumentType = "identity_card_front"
	CompanyDocumentUploadParamsDocumentTypeIdentityCardBack               CompanyDocumentUploadParamsDocumentType = "identity_card_back"
	CompanyDocumentUploadParamsDocumentTypeDriverLicenseFront             CompanyDocumentUploadParamsDocumentType = "driver_license_front"
	CompanyDocumentUploadParamsDocumentTypeDriverLicenseBack              CompanyDocumentUploadParamsDocumentType = "driver_license_back"
	CompanyDocumentUploadParamsDocumentTypeIdentityDocument               CompanyDocumentUploadParamsDocumentType = "identity_document"
	CompanyDocumentUploadParamsDocumentTypeDriverLicense                  CompanyDocumentUploadParamsDocumentType = "driver_license"
	CompanyDocumentUploadParamsDocumentTypePassport                       CompanyDocumentUploadParamsDocumentType = "passport"
	CompanyDocumentUploadParamsDocumentTypeTax                            CompanyDocumentUploadParamsDocumentType = "tax"
	CompanyDocumentUploadParamsDocumentTypeCertificateOfIncorporation     CompanyDocumentUploadParamsDocumentType = "certificate_of_incorporation"
	CompanyDocumentUploadParamsDocumentTypeCertificateOfGoodStanding      CompanyDocumentUploadParamsDocumentType = "certificate_of_good_standing"
	CompanyDocumentUploadParamsDocumentTypeLcbFtLabAmlPolicies            CompanyDocumentUploadParamsDocumentType = "lcb_ft_lab_aml_policies"
	CompanyDocumentUploadParamsDocumentTypeNiuEntreprise                  CompanyDocumentUploadParamsDocumentType = "niu_entreprise"
	CompanyDocumentUploadParamsDocumentTypeFinancialStatements            CompanyDocumentUploadParamsDocumentType = "financial_statements"
	CompanyDocumentUploadParamsDocumentTypeRccm                           CompanyDocumentUploadParamsDocumentType = "rccm"
	CompanyDocumentUploadParamsDocumentTypeProofOfSourceFunds             CompanyDocumentUploadParamsDocumentType = "proof_of_source_funds"
	CompanyDocumentUploadParamsDocumentTypeOrganizationalChart            CompanyDocumentUploadParamsDocumentType = "organizational_chart"
	CompanyDocumentUploadParamsDocumentTypeRiskPolicies                   CompanyDocumentUploadParamsDocumentType = "risk_policies"
)
