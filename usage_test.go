// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dataleonlabs_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/dataleonlabs-go"
	"github.com/stainless-sdks/dataleonlabs-go/internal/testutil"
	"github.com/stainless-sdks/dataleonlabs-go/option"
)

func TestUsage(t *testing.T) {
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
	companies, err := client.Companies.List(context.TODO(), dataleonlabs.CompanyListParams{})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", companies)
}
