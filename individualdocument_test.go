// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dataleonlabs_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/stainless-sdks/dataleonlabs-go"
	"github.com/stainless-sdks/dataleonlabs-go/internal/testutil"
	"github.com/stainless-sdks/dataleonlabs-go/option"
)

func TestIndividualDocumentList(t *testing.T) {
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
	_, err := client.Individuals.Documents.List(context.TODO(), "individual_id")
	if err != nil {
		var apierr *dataleonlabs.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIndividualDocumentUploadWithOptionalParams(t *testing.T) {
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
	_, err := client.Individuals.Documents.Upload(
		context.TODO(),
		"individual_id",
		dataleonlabs.IndividualDocumentUploadParams{
			DocumentType: dataleonlabs.IndividualDocumentUploadParamsDocumentTypeLiasseFiscale,
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
