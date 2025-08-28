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

	"github.com/dataleonlabs/dataleonlabs-go/internal/apiform"
	"github.com/dataleonlabs/dataleonlabs-go/internal/requestconfig"
	"github.com/dataleonlabs/dataleonlabs-go/option"
	"github.com/dataleonlabs/dataleonlabs-go/packages/param"
)

// IndividualDocumentService contains methods and other services that help with
// interacting with the dataleonlabs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIndividualDocumentService] method instead.
type IndividualDocumentService struct {
	Options []option.RequestOption
}

// NewIndividualDocumentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewIndividualDocumentService(opts ...option.RequestOption) (r IndividualDocumentService) {
	r = IndividualDocumentService{}
	r.Options = opts
	return
}

// Get documents to an individuals
func (r *IndividualDocumentService) List(ctx context.Context, individualID string, opts ...option.RequestOption) (res *DocumentResponse, err error) {
	opts = append(r.Options[:], opts...)
	if individualID == "" {
		err = errors.New("missing required individual_id parameter")
		return
	}
	path := fmt.Sprintf("individuals/%s/documents", individualID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Upload documents to an individual
func (r *IndividualDocumentService) Upload(ctx context.Context, individualID string, body IndividualDocumentUploadParams, opts ...option.RequestOption) (res *GenericDocument, err error) {
	opts = append(r.Options[:], opts...)
	if individualID == "" {
		err = errors.New("missing required individual_id parameter")
		return
	}
	path := fmt.Sprintf("individuals/%s/documents", individualID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type IndividualDocumentUploadParams struct {
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
	DocumentType IndividualDocumentUploadParamsDocumentType `json:"document_type,omitzero,required"`
	// URL of the file to upload (either `file` or `url` is required)
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	// File to upload (required)
	File io.Reader `json:"file,omitzero" format:"binary"`
	paramObj
}

func (r IndividualDocumentUploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
type IndividualDocumentUploadParamsDocumentType string

const (
	IndividualDocumentUploadParamsDocumentTypeLiasseFiscale                  IndividualDocumentUploadParamsDocumentType = "liasse_fiscale"
	IndividualDocumentUploadParamsDocumentTypeAmortisedLoanSchedule          IndividualDocumentUploadParamsDocumentType = "amortised_loan_schedule"
	IndividualDocumentUploadParamsDocumentTypeInvoice                        IndividualDocumentUploadParamsDocumentType = "invoice"
	IndividualDocumentUploadParamsDocumentTypeReceipt                        IndividualDocumentUploadParamsDocumentType = "receipt"
	IndividualDocumentUploadParamsDocumentTypeCompanyStatuts                 IndividualDocumentUploadParamsDocumentType = "company_statuts"
	IndividualDocumentUploadParamsDocumentTypeRegistrationCompanyCertificate IndividualDocumentUploadParamsDocumentType = "registration_company_certificate"
	IndividualDocumentUploadParamsDocumentTypeKbis                           IndividualDocumentUploadParamsDocumentType = "kbis"
	IndividualDocumentUploadParamsDocumentTypeRib                            IndividualDocumentUploadParamsDocumentType = "rib"
	IndividualDocumentUploadParamsDocumentTypeLivretFamille                  IndividualDocumentUploadParamsDocumentType = "livret_famille"
	IndividualDocumentUploadParamsDocumentTypeBirthCertificate               IndividualDocumentUploadParamsDocumentType = "birth_certificate"
	IndividualDocumentUploadParamsDocumentTypePayslip                        IndividualDocumentUploadParamsDocumentType = "payslip"
	IndividualDocumentUploadParamsDocumentTypeSocialSecurityCard             IndividualDocumentUploadParamsDocumentType = "social_security_card"
	IndividualDocumentUploadParamsDocumentTypeVehicleRegistrationCertificate IndividualDocumentUploadParamsDocumentType = "vehicle_registration_certificate"
	IndividualDocumentUploadParamsDocumentTypeCarteGrise                     IndividualDocumentUploadParamsDocumentType = "carte_grise"
	IndividualDocumentUploadParamsDocumentTypeCriminalRecordExtract          IndividualDocumentUploadParamsDocumentType = "criminal_record_extract"
	IndividualDocumentUploadParamsDocumentTypeProofOfAddress                 IndividualDocumentUploadParamsDocumentType = "proof_of_address"
	IndividualDocumentUploadParamsDocumentTypeIdentityCardFront              IndividualDocumentUploadParamsDocumentType = "identity_card_front"
	IndividualDocumentUploadParamsDocumentTypeIdentityCardBack               IndividualDocumentUploadParamsDocumentType = "identity_card_back"
	IndividualDocumentUploadParamsDocumentTypeDriverLicenseFront             IndividualDocumentUploadParamsDocumentType = "driver_license_front"
	IndividualDocumentUploadParamsDocumentTypeDriverLicenseBack              IndividualDocumentUploadParamsDocumentType = "driver_license_back"
	IndividualDocumentUploadParamsDocumentTypeIdentityDocument               IndividualDocumentUploadParamsDocumentType = "identity_document"
	IndividualDocumentUploadParamsDocumentTypeDriverLicense                  IndividualDocumentUploadParamsDocumentType = "driver_license"
	IndividualDocumentUploadParamsDocumentTypePassport                       IndividualDocumentUploadParamsDocumentType = "passport"
	IndividualDocumentUploadParamsDocumentTypeTax                            IndividualDocumentUploadParamsDocumentType = "tax"
	IndividualDocumentUploadParamsDocumentTypeCertificateOfIncorporation     IndividualDocumentUploadParamsDocumentType = "certificate_of_incorporation"
	IndividualDocumentUploadParamsDocumentTypeCertificateOfGoodStanding      IndividualDocumentUploadParamsDocumentType = "certificate_of_good_standing"
	IndividualDocumentUploadParamsDocumentTypeLcbFtLabAmlPolicies            IndividualDocumentUploadParamsDocumentType = "lcb_ft_lab_aml_policies"
	IndividualDocumentUploadParamsDocumentTypeNiuEntreprise                  IndividualDocumentUploadParamsDocumentType = "niu_entreprise"
	IndividualDocumentUploadParamsDocumentTypeFinancialStatements            IndividualDocumentUploadParamsDocumentType = "financial_statements"
	IndividualDocumentUploadParamsDocumentTypeRccm                           IndividualDocumentUploadParamsDocumentType = "rccm"
	IndividualDocumentUploadParamsDocumentTypeProofOfSourceFunds             IndividualDocumentUploadParamsDocumentType = "proof_of_source_funds"
	IndividualDocumentUploadParamsDocumentTypeOrganizationalChart            IndividualDocumentUploadParamsDocumentType = "organizational_chart"
	IndividualDocumentUploadParamsDocumentTypeRiskPolicies                   IndividualDocumentUploadParamsDocumentType = "risk_policies"
)
