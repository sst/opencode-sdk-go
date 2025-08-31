// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/opencode-go"
	"github.com/stainless-sdks/opencode-go/internal/testutil"
	"github.com/stainless-sdks/opencode-go/option"
)

func TestAuthSetCredentials(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Auth.SetCredentials(
		context.TODO(),
		"id",
		opencode.AuthSetCredentialsParams{
			OfOAuth: &opencode.AuthSetCredentialsParamsBodyOAuth{
				Access:  "access",
				Expires: 0,
				Refresh: "refresh",
			},
		},
	)
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
