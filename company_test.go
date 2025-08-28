// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dataleonlabs_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/dataleonlabs/dataleonlabs-go"
	"github.com/dataleonlabs/dataleonlabs-go/internal/testutil"
	"github.com/dataleonlabs/dataleonlabs-go/option"
)

func TestCompanyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Companies.New(context.TODO(), dataleonlabs.CompanyNewParams{
		Company: dataleonlabs.CompanyNewParamsCompany{
			Name:                         "ACME Corp",
			Address:                      dataleonlabs.String("123 rue Exemple, Paris"),
			CommercialName:               dataleonlabs.String("ACME"),
			Country:                      dataleonlabs.String("FR"),
			Email:                        dataleonlabs.String("info@acme.fr"),
			EmployerIdentificationNumber: dataleonlabs.String("EIN123456"),
			LegalForm:                    dataleonlabs.String("SARL"),
			PhoneNumber:                  dataleonlabs.String("+33 1 23 45 67 89"),
			RegistrationDate:             dataleonlabs.String("2010-05-15"),
			RegistrationID:               dataleonlabs.String("RCS123456"),
			ShareCapital:                 dataleonlabs.String("100000"),
			Status:                       dataleonlabs.String("active"),
			TaxIdentificationNumber:      dataleonlabs.String("FR123456789"),
			Type:                         dataleonlabs.String("main"),
			WebsiteURL:                   dataleonlabs.String("https://acme.fr"),
		},
		WorkspaceID: "wk_123",
		SourceID:    dataleonlabs.String("ID54410069066"),
		TechnicalData: dataleonlabs.CompanyNewParamsTechnicalData{
			ActiveAmlSuspicions:     dataleonlabs.Bool(false),
			CallbackURL:             dataleonlabs.String("https://example.com/callback"),
			CallbackURLNotification: dataleonlabs.String("https://example.com/notify"),
			Language:                dataleonlabs.String("fra"),
			RawData:                 dataleonlabs.Bool(true),
		},
	})
	if err != nil {
		var apierr *dataleonlabs.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCompanyGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Companies.Get(
		context.TODO(),
		"company_id",
		dataleonlabs.CompanyGetParams{
			Document: dataleonlabs.Bool(true),
			Scope:    dataleonlabs.String("scope"),
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

func TestCompanyUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Companies.Update(
		context.TODO(),
		"company_id",
		dataleonlabs.CompanyUpdateParams{
			Company: dataleonlabs.CompanyUpdateParamsCompany{
				Name:                         "ACME Corp",
				Address:                      dataleonlabs.String("123 rue Exemple, Paris"),
				CommercialName:               dataleonlabs.String("ACME"),
				Country:                      dataleonlabs.String("FR"),
				Email:                        dataleonlabs.String("info@acme.fr"),
				EmployerIdentificationNumber: dataleonlabs.String("EIN123456"),
				LegalForm:                    dataleonlabs.String("SARL"),
				PhoneNumber:                  dataleonlabs.String("+33 1 23 45 67 89"),
				RegistrationDate:             dataleonlabs.String("2010-05-15"),
				RegistrationID:               dataleonlabs.String("RCS123456"),
				ShareCapital:                 dataleonlabs.String("100000"),
				Status:                       dataleonlabs.String("active"),
				TaxIdentificationNumber:      dataleonlabs.String("FR123456789"),
				Type:                         dataleonlabs.String("main"),
				WebsiteURL:                   dataleonlabs.String("https://acme.fr"),
			},
			WorkspaceID: "wk_123",
			SourceID:    dataleonlabs.String("ID54410069066"),
			TechnicalData: dataleonlabs.CompanyUpdateParamsTechnicalData{
				ActiveAmlSuspicions:     dataleonlabs.Bool(false),
				CallbackURL:             dataleonlabs.String("https://example.com/callback"),
				CallbackURLNotification: dataleonlabs.String("https://example.com/notify"),
				Language:                dataleonlabs.String("fra"),
				RawData:                 dataleonlabs.Bool(true),
			},
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

func TestCompanyListWithOptionalParams(t *testing.T) {
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
	_, err := client.Companies.List(context.TODO(), dataleonlabs.CompanyListParams{
		EndDate:     dataleonlabs.Time(time.Now()),
		Limit:       dataleonlabs.Int(1),
		Offset:      dataleonlabs.Int(0),
		SourceID:    dataleonlabs.String("source_id"),
		StartDate:   dataleonlabs.Time(time.Now()),
		State:       dataleonlabs.CompanyListParamsStateVoid,
		Status:      dataleonlabs.CompanyListParamsStatusRejected,
		WorkspaceID: dataleonlabs.String("workspace_id"),
	})
	if err != nil {
		var apierr *dataleonlabs.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCompanyDelete(t *testing.T) {
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
	err := client.Companies.Delete(context.TODO(), "company_id")
	if err != nil {
		var apierr *dataleonlabs.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
