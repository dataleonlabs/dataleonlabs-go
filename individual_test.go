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

func TestIndividualNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Individuals.New(context.TODO(), dataleonlabs.IndividualNewParams{
		WorkspaceID: "wk_123",
		Person: dataleonlabs.IndividualNewParamsPerson{
			Birthday:    dataleonlabs.String("15/05/1985"),
			Email:       dataleonlabs.String("john.doe@example.com"),
			FirstName:   dataleonlabs.String("John"),
			Gender:      "M",
			LastName:    dataleonlabs.String("Doe"),
			MaidenName:  dataleonlabs.String("John Doe"),
			Nationality: dataleonlabs.String("FRA"),
			PhoneNumber: dataleonlabs.String("+33 1 23 45 67 89"),
		},
		SourceID: dataleonlabs.String("ID54410069066"),
		TechnicalData: dataleonlabs.IndividualNewParamsTechnicalData{
			ActiveAmlSuspicions:         dataleonlabs.Bool(false),
			CallbackURL:                 dataleonlabs.String("https://example.com/callback"),
			CallbackURLNotification:     dataleonlabs.String("https://example.com/notify"),
			FilteringScoreAmlSuspicions: dataleonlabs.Float(0.75),
			Language:                    dataleonlabs.String("fra"),
			RawData:                     dataleonlabs.Bool(true),
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

func TestIndividualGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Individuals.Get(
		context.TODO(),
		"individual_id",
		dataleonlabs.IndividualGetParams{
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

func TestIndividualUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Individuals.Update(
		context.TODO(),
		"individual_id",
		dataleonlabs.IndividualUpdateParams{
			WorkspaceID: "wk_123",
			Person: dataleonlabs.IndividualUpdateParamsPerson{
				Birthday:    dataleonlabs.String("15/05/1985"),
				Email:       dataleonlabs.String("john.doe@example.com"),
				FirstName:   dataleonlabs.String("John"),
				Gender:      "M",
				LastName:    dataleonlabs.String("Doe"),
				MaidenName:  dataleonlabs.String("John Doe"),
				Nationality: dataleonlabs.String("FRA"),
				PhoneNumber: dataleonlabs.String("+33 1 23 45 67 89"),
			},
			SourceID: dataleonlabs.String("ID54410069066"),
			TechnicalData: dataleonlabs.IndividualUpdateParamsTechnicalData{
				ActiveAmlSuspicions:         dataleonlabs.Bool(false),
				CallbackURL:                 dataleonlabs.String("https://example.com/callback"),
				CallbackURLNotification:     dataleonlabs.String("https://example.com/notify"),
				FilteringScoreAmlSuspicions: dataleonlabs.Float(0.75),
				Language:                    dataleonlabs.String("fra"),
				RawData:                     dataleonlabs.Bool(true),
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

func TestIndividualListWithOptionalParams(t *testing.T) {
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
	_, err := client.Individuals.List(context.TODO(), dataleonlabs.IndividualListParams{
		EndDate:     dataleonlabs.Time(time.Now()),
		Limit:       dataleonlabs.Int(1),
		Offset:      dataleonlabs.Int(0),
		SourceID:    dataleonlabs.String("source_id"),
		StartDate:   dataleonlabs.Time(time.Now()),
		State:       dataleonlabs.IndividualListParamsStateVoid,
		Status:      dataleonlabs.IndividualListParamsStatusRejected,
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

func TestIndividualDelete(t *testing.T) {
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
	err := client.Individuals.Delete(context.TODO(), "individual_id")
	if err != nil {
		var apierr *dataleonlabs.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
