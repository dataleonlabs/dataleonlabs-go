// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dataleonlabs_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/dataleonlabs/dataleonlabs-go"
	"github.com/dataleonlabs/dataleonlabs-go/internal/testutil"
	"github.com/dataleonlabs/dataleonlabs-go/option"
)

func TestCompanyDocumentList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := dataleonlabs.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Companies.Documents.List(context.TODO(), "company_id")
	if err != nil {
		var apierr *dataleonlabs.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCompanyDocumentUploadWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := dataleonlabs.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Companies.Documents.Upload(
		context.TODO(),
		"company_id",
		dataleonlabs.CompanyDocumentUploadParams{
			DocumentType: dataleonlabs.CompanyDocumentUploadParamsDocumentTypeLiasseFiscale,
			File:         io.Reader(bytes.NewBuffer([]byte("some file contents"))),
			URL:          dataleonlabs.String("https://example.com/sample.pdf"),
		},
	)
	if err != nil {
		var apierr *dataleonlabs.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
